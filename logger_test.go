package fiberlogrus

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	t.Parallel()
	app := fiber.New()

	s := ""
	buf := bytes.NewBufferString(s)

	logger := logrus.New()
	logger.SetOutput(buf)

	app.Use(New(
		Config{
			Logger: logger,
			Tags: []string{
				TagMethod,
				TagStatus,
				AttachKeyTag(TagLocals, "loc"),
				AttachKeyTag(TagRespHeader, "custom-header"),
			},
		}))

	app.Get("/", func(c *fiber.Ctx) error {
		c.Append("custom-header", "custom-header-value")
		c.Locals("loc", "val")
		return c.SendString("random string")
	})
	resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/", nil))
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)
	require.Contains(t, buf.String(), "method=GET")
	require.Contains(t, buf.String(), "status=200")
	require.Contains(t, buf.String(), "respHeader=custom-header-value")
	require.Contains(t, buf.String(), "locals=val")
}
