package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"golang.org/x/net/context"
)

var (
	containerId string
)

func init() {
	flag.StringVar(&containerId, "-filter", "", "Filter containers by name")
	flag.StringVar(&containerId, "f", "", "Filter containers by name")
	flag.Parse()
}

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

	fmt.Println()
}

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	sort.Slice(containers, func(i, j int) bool {
		return containers[i].Names[0] < containers[j].Names[0]
	})

	if containerId == "" {
		for _, container := range containers {
			printContainerInfo(container)
		}
	} else {
		var found int

		for _, container := range containers {
			if strings.Contains(container.Names[0], containerId) {
				found++
				printContainerInfo(container)
			}
		}

		if found == 0 {
			fmt.Println("No containers found for pattern: ", containerId)
		}
	}
}
