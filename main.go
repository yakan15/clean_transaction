package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/yakan15/clean-transaction/handler"
	"github.com/yakan15/clean-transaction/infra"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i any) error {
	return v.validator.Struct(i)
}

func newRouter() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Validator = &Validator{validator: validator.New()}
	return e
}

func main() {
	r := newRouter()
	// api := r.Group("/api")
	r.GET("/test", func(c echo.Context) error { return c.NoContent(http.StatusNoContent) })
	d := infra.NewDb()
	infra.AutoMigrate(d)

	h := handler.NewHandler()
	h.Register(r)
	r.Logger.Fatal(r.Start("127.0.0.1:8080"))
}
