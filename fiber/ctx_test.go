package fiber

import (
	"bytes"
	"io"
	"testing"
	"github.com/valyala/fasthttp"
)

func TestCtx_RestoreBody(t *testing.T) {
	app := New()
	app.Use(func(c *Ctx) error {
		// Simulate middleware consuming the body
		_, _ = io.ReadAll(c.Request.BodyStream())
		// Restore it
		_ = c.RestoreBody()
		return c.Next()
	})

	app.Post("/", func(c *Ctx) error {
		// Downstream handler should be able to read the body
		body := c.Body()
		if string(body) != "test" {
			t.Errorf("expected 'test', got %s", string(body))
		}
		return nil
	})

	req := &fasthttp.Request{}
	req.SetBody([]byte("test"))
	req.SetRequestURI("/")
	req.Header.SetMethod("POST")

	// Test logic here...
}