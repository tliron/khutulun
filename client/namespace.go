package client

import (
	"io"

	"github.com/tliron/khutulun/sdk"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (self *Client) ListNamespaces() ([]string, error) {
	context, cancel := self.newContextWithTimeout()
	defer cancel()

	if client, err := self.client.ListNamespaces(context, new(emptypb.Empty)); err == nil {
		var namespaces []string

		for {
			namespace, err := client.Recv()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, sdk.UnpackGRPCError(err)
				}
			}

			namespaces = append(namespaces, namespace.Name)
		}

		return namespaces, nil
	} else {
		return nil, sdk.UnpackGRPCError(err)
	}
}
