// Code generated by thriftrw-plugin-yarpc
// @generated

package readonlystoreclient

import (
	"context"
	"go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc/v2"
	"go.uber.org/yarpc/v2/yarpcthrift"
	"go.uber.org/yarpc/v2/yarpcthrift/thriftrw-plugin-yarpc2/internal/tests/atomic"
	"go.uber.org/yarpc/v2/yarpcthrift/thriftrw-plugin-yarpc2/internal/tests/common/baseserviceclient"
)

// Interface is a client for the ReadOnlyStore service.
type Interface interface {
	baseserviceclient.Interface

	Integer(
		ctx context.Context,
		Key *string,
		opts ...yarpc.CallOption,
	) (int64, error)
}

// New builds a new client for the ReadOnlyStore service.
//
//  yarpcClient, err := yarpc.Provider.Client("readonlystore")
//  if err != nil {
//	  return err
//  }
// 	client := readonlystoreclient.New(yarpcClient)
func New(c yarpc.Client, opts ...yarpcthrift.ClientOption) Interface {
	return client{
		c: yarpcthrift.New(
			c,
			"ReadOnlyStore",
			opts...),
		Interface: baseserviceclient.New(c, opts...),
	}
}

type client struct {
	baseserviceclient.Interface

	c yarpcthrift.Client
}

func (c client) Integer(
	ctx context.Context,
	_Key *string,
	opts ...yarpc.CallOption,
) (success int64, err error) {

	args := atomic.ReadOnlyStore_Integer_Helper.Args(_Key)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result atomic.ReadOnlyStore_Integer_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = atomic.ReadOnlyStore_Integer_Helper.UnwrapResponse(&result)
	return
}