package main

import (
	"fmt"
	"strings"

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
		fmt.Println()
	}
}
