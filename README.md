# mtgo middleware: YOUR_MIDDLEWARE_NAME

Short description of what this middleware does.

## Install

```bash
go get github.com/mtgo-labs/middlewares/YOUR_MIDDLEWARE_NAME
```

## Getting Started

1. Copy this template to `middlewares/your_middleware_name/`
2. Replace `YOUR_MIDDLEWARE_NAME` in `go.mod`, `middleware.go`, `middleware_test.go`, and `README.md` with your middleware name (lowercase, no hyphens/underscores)
3. Update `go.mod` module path to `github.com/mtgo-labs/middlewares/your_middleware_name`
4. Implement the middleware logic in the `Middleware()` method

## Usage

```go
import (
    tg "github.com/mtgo-labs/mtgo/telegram"
    yourmw "github.com/mtgo-labs/middlewares/YOUR_MIDDLEWARE_NAME"
)

func main() {
    client, _ := tg.NewClient(apiID, apiHash, &tg.Config{
        BotToken: botToken,
    })

    mw := yourmw.New(yourmw.Config{})
    client.UseInvokerMiddleware(mw.Middleware())

    // your handlers...
    client.Start()
}
```

## Architecture

mtgo has two middleware levels:

| Level | Type | Intercepts | Use case |
|-------|------|------------|----------|
| Invoker | `UseInvokerMiddleware` | RPC calls (API calls to Telegram) | Rate limiting, flood wait, logging, metrics |
| Handler | `UseMiddleware` | Update dispatch (incoming messages) | Auth, i18n, conversation state |

This template implements an **invoker middleware** — it wraps outgoing RPC calls. No handler changes are required.

## License

MIT
