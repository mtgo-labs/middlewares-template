package YOUR_MIDDLEWARE_NAME

import (
	"context"

	"github.com/mtgo-labs/mtgo/tg"
)

// Config holds middleware configuration.
type Config struct {
	// Add your configuration fields here.
}

// Middleware is an RPC invoker middleware that wraps each Telegram API call.
// It intercepts outgoing RPC calls and can inspect, modify, or retry them.
type Middleware struct {
	config Config
}

// New creates a new middleware instance with the given configuration.
func New(config Config) *Middleware {
	return &Middleware{config: config}
}

// Middleware returns a tg.InvokerMiddleware function for UseInvokerMiddleware.
// This intercepts ALL outgoing RPC calls — no handler changes required.
//
// Usage:
//
//	mw := yourmiddleware.New(yourmiddleware.Config{})
//	client.UseInvokerMiddleware(mw.Middleware())
func (m *Middleware) Middleware() func(next tg.Invoker) tg.Invoker {
	return func(next tg.Invoker) tg.Invoker {
		return tg.InvokerFunc(func(ctx context.Context, input tg.TLObject, decode func(*tg.Reader) (tg.TLObject, error)) (tg.TLObject, error) {
			// Pre-processing: inspect input, add logging, etc.

			// Call the next invoker in the chain.
			result, err := next.RPCInvoke(ctx, input, decode)

			// Post-processing: inspect result/error, handle retries, etc.

			return result, err
		})
	}
}
