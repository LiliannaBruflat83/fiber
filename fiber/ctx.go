package fiber

import (
	"io"
	"github.com/valyala/fasthttp"
)

// RestoreBody reads the current request body stream and caches it into the request.
// This is useful when a middleware has consumed the body stream and it needs to be available for downstream handlers.
func (c *Ctx) RestoreBody() error {
	body, err := io.ReadAll(c.Request.BodyStream())
	if err != nil {
		return err
	}
	c.Request.SetBody(body)
	return nil
}