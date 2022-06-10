package sdk

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tliron/kutil/ard"
	cloutpkg "github.com/tliron/puccini/clout"
)

//
// Container
//

type Container struct {
	Host            string
	Name            string
	Reference       string
	CreateArguments []string
	Ports           []Port

	vertexID       string
	capabilityName string
}

type Port struct {
	External int64
	Internal int64
	Protocol string
}

func (self *Container) Find(clout *cloutpkg.Clout) (*cloutpkg.Vertex, any, error) {
	if vertex, ok := clout.Vertexes[self.vertexID]; ok {
		if capabilities, ok := ard.NewNode(vertex.Properties).Get("capabilities").StringMap(); ok {
			if capability, ok := capabilities[self.capabilityName]; ok {
				return vertex, capability, nil
			} else {
				return nil, nil, fmt.Errorf("vertex %s has no capability: %s", self.vertexID, self.capabilityName)
			}
		} else {
			return nil, nil, fmt.Errorf("vertex has no capabilities: %s", self.vertexID)
		}
	} else {
		return nil, nil, fmt.Errorf("vertex not found: %s", self.vertexID)
	}
}

func GetContainerPorts(capability any) []Port {
	var ports []Port
	capabilityAttributes, _ := ard.NewNode(capability).Get("attributes").StringMap()
	if ports_, ok := ard.NewNode(capabilityAttributes).Get("ports").List(); ok {
		for _, port := range ports_ {
			external, _ := ard.NewNode(port).Get("external").NumberAsInteger()
			internal, _ := ard.NewNode(port).Get("internal").NumberAsInteger()
			protocol, _ := ard.NewNode(port).Get("protocol").String()
			ports = append(ports, Port{
				External: external,
				Internal: internal,
				Protocol: protocol,
			})
		}
	}
	return ports
}

func GetContainers(vertex *cloutpkg.Vertex, capabilityName string, capability any) []*Container {
	var containers []*Container

	instances, _ := ard.NewNode(vertex.Properties).Get("attributes").Get("instances").List()
	for index, instance := range instances {
		container := Container{
			vertexID:       vertex.ID,
			capabilityName: capabilityName,
		}

		capabilityProperties, _ := ard.NewNode(capability).Get("properties").StringMap()

		container.Host, _ = ard.NewNode(instance).Get("host").String()
		var ok bool
		if container.Name, ok = ard.NewNode(capabilityProperties).Get("name").String(); !ok {
			container.Name, _ = ard.NewNode(vertex.Properties).Get("name").String()
		}
		container.Name = fmt.Sprintf("%s-%d", container.Name, index)
		container.Reference, _ = ard.NewNode(capabilityProperties).Get("image").Get("reference").String()
		if container.Reference == "" {
			host, _ := ard.NewNode(capabilityProperties).Get("image").Get("host").String()
			port, _ := ard.NewNode(capabilityProperties).Get("image").Get("port").NumberAsInteger()
			repository, _ := ard.NewNode(capabilityProperties).Get("image").Get("repository").String()
			image, _ := ard.NewNode(capabilityProperties).Get("image").Get("image").String()
			tag, _ := ard.NewNode(capabilityProperties).Get("image").Get("tag").String()
			digestAlgorithm, _ := ard.NewNode(capabilityProperties).Get("image").Get("digest-algorithm").String()
			digestHex, _ := ard.NewNode(capabilityProperties).Get("image").Get("digest-hex").String()
			if image != "" {
				container.Reference = formatImageReference(host, int(port), repository, image, tag, digestAlgorithm, digestHex)
			}
		}
		container.CreateArguments, _ = ard.NewNode(capabilityProperties).Get("create-arguments").StringList()

		containers = append(containers, &container)
	}

	return containers
}

func GetVertexContainers(vertex *cloutpkg.Vertex) []*Container {
	var containers []*Container
	if capabilities, ok := ard.NewNode(vertex.Properties).Get("capabilities").StringMap(); ok {
		for capabilityName, capability := range capabilities {
			if types, ok := ard.NewNode(capability).Get("types").StringMap(); ok {
				if _, ok := types["cloud.puccini.khutulun::Container"]; ok {
					containers = append(containers, GetContainers(vertex, capabilityName, capability)...)
				}
			}
		}

		for _, capability := range capabilities {
			if types, ok := ard.NewNode(capability).Get("types").StringMap(); ok {
				if _, ok := types["cloud.puccini.khutulun::ContainerConnectable"]; ok {
					ports := GetContainerPorts(capability)
					for _, container := range containers {
						container.Ports = ports
					}
				}
			}
		}
	}
	return containers
}

func GetCloutContainers(clout *cloutpkg.Clout) []*Container {
	var containers []*Container
	for _, vertex := range clout.Vertexes {
		containers = append(containers, GetVertexContainers(vertex)...)
	}
	return containers
}

func formatImageReference(host string, port int, repository string, image string, tag string, digestAlgorithm string, digestHex string) string {
	// [host[:port]/][repository/]image[:tag][@digest-algorithm:digest-hex]
	var s strings.Builder
	if host != "" {
		s.WriteString(host)
		if port != 0 {
			s.WriteRune(':')
			s.WriteString(strconv.Itoa(port))
		}
		s.WriteRune('/')
	}
	if repository != "" {
		s.WriteString(repository)
		s.WriteRune('/')
	}
	s.WriteString(image)
	if tag != "" {
		s.WriteRune(':')
		s.WriteString(tag)
	}
	if digestAlgorithm != "" {
		s.WriteRune('@')
		s.WriteString(digestAlgorithm)
		if digestHex != "" {
			s.WriteRune(':')
			s.WriteString(digestHex)
		}
	}
	return s.String()
}
