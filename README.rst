dcti - Docker container info
============================

Pet project to toy with `Go`_ and the `Docker API`_.

.. _Go: https://golang.org
.. _Docker API: https://docs.docker.com/develop/sdk/

Resources
---------

Docker API
~~~~~~~~~~

- https://docs.docker.com/develop/sdk/examples/

Reference:

- https://godoc.org/golang.org/x/net/context
- https://godoc.org/github.com/docker/docker/client
- https://godoc.org/github.com/docker/docker/client#Client.ClientVersion
- https://godoc.org/github.com/docker/docker/client#Client.ContainerList
- https://godoc.org/github.com/docker/docker/api/types#Container
- https://godoc.org/github.com/docker/docker/api/types#ContainerListOptions
- https://godoc.org/github.com/docker/docker/api/types#MountPoint
- https://godoc.org/github.com/docker/docker/api/types/network#EndpointSettings
- https://godoc.org/github.com/docker/docker/api/types#SummaryNetworkSettings

Sources:

- https://github.com/moby/moby/tree/master/container
- https://github.com/moby/moby/tree/master/api/types

Go
~~

- https://golang.org/pkg/flag/
- https://golang.org/pkg/sort/
- https://golang.org/pkg/sort/#Slice
- https://golang.org/pkg/strings/

Articles:

- https://gobyexample.com/command-line-flags

Threads:

- https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field
- https://stackoverflow.com/questions/23466497/how-to-truncate-a-string-in-a-golang-template

Inspiration
~~~~~~~~~~~

- https://github.com/robphoenix/trawl
- https://github.com/genuinetools/weather
