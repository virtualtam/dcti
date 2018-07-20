package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"golang.org/x/net/context"
)

func printContainerInfo(container types.Container) {
	fmt.Println("===", strings.TrimLeft(container.Names[0], "/"), "===")
	fmt.Println("Image:", container.Image)

	for endpoint, settings := range container.NetworkSettings.Networks {
		fmt.Printf("\nNetwork '%s':\n", endpoint)
		fmt.Println("  Gateway:", settings.Gateway)
		fmt.Println("  Address:", settings.IPAddress)
		fmt.Println("  MAC:    ", settings.MacAddress)
	}

	if len(container.Mounts) > 0 {
		fmt.Println("\nMounts:")
		for _, mount := range container.Mounts {
			fmt.Printf("  %s:\n", mount.Destination)
			fmt.Println("    Type:  ", mount.Type)
			fmt.Println("    Name:  ", mount.Name)
			fmt.Println("    Source:", mount.Source)
			fmt.Println("    Mode:  ", mount.Mode)
		}
	}
}

func renderTpl(writer *tabwriter.Writer, templateString string, data interface{}) {
	tpl := template.New("tpl")
	tpl.Parse(templateString)
	tpl.Execute(writer, data)
	writer.Flush()
}

func printTpl(container types.Container) {
	const (
		info = "Names\tImage\n" +
			"{{range $n := .Names}}{{$n}}{{end}}\t{{.Image}}\n\n"
		net = "Gateway\tIP\tMAC\n" +
			"{{range .Networks}}" +
			"{{.Gateway}}\t{{.IPAddress}}\t{{.MacAddress}}" +
			"{{end}}\n\n"
		mount = "Type\tName\tSource\tMode\n" +
			"{{range .Mounts}}" +
			"{{.Type}}\t{{.Name}}\t{{.Source}}\t{{.Mode}}" +
			"{{end}}\n"
	)
	writer := tabwriter.NewWriter(os.Stdout, 8, 8, 8, ' ', 0)

	fmt.Println("===", strings.TrimLeft(container.Names[0], "/"), "===")

	renderTpl(writer, info, container)
	renderTpl(writer, net, container.NetworkSettings)
	renderTpl(writer, mount, container)
}

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		printContainerInfo(container)
		printTpl(container)
		fmt.Println()
	}
}
