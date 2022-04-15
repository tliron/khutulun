package client

import (
	"io"

	"github.com/tliron/khutulun/api"
)

type Resource struct {
	Namespace string `json:"namespace" yaml:"namespace"`
	Service   string `json:"service" yaml:"service"`
	Type      string `json:"type" yaml:"type"`
	Name      string `json:"name" yaml:"name"`
}

func (self *Client) ListResources(namespace string, serviceName string, type_ string) ([]Resource, error) {
	context, cancel := self.newContextWithTimeout()
	defer cancel()

	listResources := api.ListResources{
		Service: &api.ServiceIdentifier{
			Namespace: namespace,
			Name:      serviceName,
		},
		Type: type_,
	}

	if client, err := self.client.ListResources(context, &listResources); err == nil {
		var resources []Resource

		for {
			identifier, err := client.Recv()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, err
				}
			}

			resources = append(resources, Resource{
				Namespace: identifier.Service.Namespace,
				Service:   identifier.Service.Name,
				Type:      identifier.Type,
				Name:      identifier.Name,
			})
		}

		return resources, nil
	} else {
		return nil, err
	}
}
