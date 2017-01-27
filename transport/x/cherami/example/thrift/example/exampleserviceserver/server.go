// Code generated by thriftrw-plugin-yarpc
// @generated

package exampleserviceserver

import (
	"context"

	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/transport/x/cherami/example/thrift/example"
)

// Interface is the server-side interface for the ExampleService service.
type Interface interface {
	Award(
		ctx context.Context,
		Token *string,
	) error
}

// New prepares an implementation of the ExampleService service for
// registration.
//
// 	handler := ExampleServiceHandler{}
// 	dispatcher.Register(exampleserviceserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "ExampleService",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "award",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Oneway,
					Oneway: thrift.OnewayHandler(h.Award),
				},
				Signature: "Award(Token *string)",
			},
		},
	}
	return thrift.BuildProcedures(service, opts...)
}

type handler struct{ impl Interface }

func (h handler) Award(ctx context.Context, body wire.Value) error {
	var args example.ExampleService_Award_Args
	if err := args.FromWire(body); err != nil {
		return err
	}

	return h.impl.Award(ctx, args.Token)
}