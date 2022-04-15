package util

import (
	"io"

	"github.com/tliron/khutulun/api"
	"github.com/tliron/kutil/logging"
	"google.golang.org/grpc/codes"
	statuspkg "google.golang.org/grpc/status"
)

type InteractFunc func(first *api.Interaction) error

func Interact(server Interactor, interact map[string]InteractFunc) error {
	if first, err := server.Recv(); err == nil {
		if first.Start != nil {
			if len(first.Start.Identifier) == 0 {
				return statuspkg.Error(codes.InvalidArgument, "no identifier")
			}
			type_ := first.Start.Identifier[0]

			if interact_, ok := interact[type_]; ok {
				return interact_(first)
			} else {
				return statuspkg.Errorf(codes.InvalidArgument, "malformed identifier: %s", first.Start.Identifier)
			}
		} else {
			return statuspkg.Errorf(codes.InvalidArgument, "first message must contain \"start\"")
		}
	} else {
		return statuspkg.Errorf(codes.Aborted, "%s", err.Error())
	}
}

func InteractRelay(server Interactor, client Interactor, first *api.Interaction, log logging.Logger) error {
	if err := client.Send(first); err != nil {
		return err
	}

	go func() {
		for {
			if interaction, err := server.Recv(); err == nil {
				if err := client.Send(interaction); err != nil {
					log.Errorf("client send: %s", err.Error())
					return
				}
			} else {
				if err == io.EOF {
					log.Info("client closed")
				} else {
					if status, ok := statuspkg.FromError(err); ok {
						if status.Code() == codes.Canceled {
							// We're OK with canceling
							log.Infof("client canceled")
							return
						}
					}
				}
				log.Errorf("client receive: %s", err.Error())
				return
			}
		}
	}()

	for {
		if interaction, err := client.Recv(); err == nil {
			if err := server.Send(interaction); err != nil {
				return err
			}
		} else {
			if err == io.EOF {
				log.Info("server closed")
				return nil
			} else {
				return err
			}
		}
	}
}
