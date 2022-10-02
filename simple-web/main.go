package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Body map[string]any

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	renderer := new(Renderer)
	renderer.location = location
	renderer.debug = debug

	renderer.ReloadTemplates()

	return renderer
}

func (r *Renderer) ReloadTemplates() {
	r.template = template.Must(template.ParseGlob(r.location))
}

func (r *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if r.debug {
		r.ReloadTemplates()
	}

	return r.template.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = NewRenderer("./*.html", true)

	e.GET("/index", func(c echo.Context) error {
		body := Body{"message": "apa kabs"}
		return c.Render(http.StatusOK, "index.html", body)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
