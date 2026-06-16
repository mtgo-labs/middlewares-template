package YOUR_MIDDLEWARE_NAME_test

import (
	"context"
	"errors"
	"testing"

	"github.com/mtgo-labs/mtgo/tg"
	"github.com/mtgo-labs/middlewares/YOUR_MIDDLEWARE_NAME"
)

type mockInvoker struct {
	fn func(context.Context, tg.TLObject, func(*tg.Reader) (tg.TLObject, error)) (tg.TLObject, error)
}

func (m *mockInvoker) RPCInvoke(ctx context.Context, input tg.TLObject, decode func(*tg.Reader) (tg.TLObject, error)) (tg.TLObject, error) {
	return m.fn(ctx, input, decode)
}

func (m *mockInvoker) RPCInvokeRaw(_ context.Context, _ tg.TLObject) ([]byte, error) {
	return nil, nil
}

func TestNew(t *testing.T) {
	mw := YOUR_MIDDLEWARE_NAME.New(YOUR_MIDDLEWARE_NAME.Config{})
	if mw == nil {
		t.Fatal("expected non-nil middleware")
	}
}

func TestPassthrough(t *testing.T) {
	base := &mockInvoker{fn: func(_ context.Context, _ tg.TLObject, _ func(*tg.Reader) (tg.TLObject, error)) (tg.TLObject, error) {
		return nil, nil
	}}

	invoker := YOUR_MIDDLEWARE_NAME.New(YOUR_MIDDLEWARE_NAME.Config{}).Middleware()(base)
	_, err := invoker.RPCInvoke(context.Background(), nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestErrorPassthrough(t *testing.T) {
	expectedErr := errors.New("test error")
	base := &mockInvoker{fn: func(_ context.Context, _ tg.TLObject, _ func(*tg.Reader) (tg.TLObject, error)) (tg.TLObject, error) {
		return nil, expectedErr
	}}

	invoker := YOUR_MIDDLEWARE_NAME.New(YOUR_MIDDLEWARE_NAME.Config{}).Middleware()(base)
	_, err := invoker.RPCInvoke(context.Background(), nil, nil)
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected original error, got: %v", err)
	}
}
