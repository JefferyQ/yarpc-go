// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package yarpcfxmiddleware

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/config"
	yarpc "go.uber.org/yarpc/v2"
)

func TestNewInboundTransportConfig(t *testing.T) {
	cfg := strings.NewReader(`yarpc: {middleware: {inbounds: {transport: {unary: ["nop"]}}}}`)
	provider, err := config.NewYAML(config.Source(cfg))
	require.NoError(t, err)

	res, err := newInboundTransportConfig(InboundTransportConfigParams{
		Provider: provider,
	})
	require.NoError(t, err)
	assert.Equal(t, InboundTransportConfig{Unary: []string{"nop"}}, res.Config)
}

func TestNewUnaryInboundTransport(t *testing.T) {
	t.Run("duplicate registration error", func(t *testing.T) {
		_, err := newUnaryInboundTransport(
			UnaryInboundTransportParams{
				Middleware: []yarpc.UnaryInboundTransportMiddleware{
					yarpc.NopUnaryInboundTransportMiddleware,
					yarpc.NopUnaryInboundTransportMiddleware,
				},
			},
		)
		assert.EqualError(t, err, `unary inbound transport middleware "nop" was registered more than once`)
	})

	t.Run("configured middleware is not available", func(t *testing.T) {
		_, err := newUnaryInboundTransport(
			UnaryInboundTransportParams{
				Config: InboundTransportConfig{
					Unary: []string{"dne"},
				},
			},
		)
		assert.EqualError(t, err, `failed to resolve unary inbound transport middleware: "dne"`)
	})

	t.Run("successful construction", func(t *testing.T) {
		res, err := newUnaryInboundTransport(
			UnaryInboundTransportParams{
				Config: InboundTransportConfig{
					Unary: []string{"nop"},
				},
				MiddlewareLists: [][]yarpc.UnaryInboundTransportMiddleware{
					{
						yarpc.NopUnaryInboundTransportMiddleware,
					},
				},
			},
		)
		require.NoError(t, err)

		middleware := res.OrderedMiddleware
		require.Len(t, middleware, 1)
		assert.Equal(t, "nop", middleware[0].Name())
	})
}