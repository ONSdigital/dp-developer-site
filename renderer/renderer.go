package renderer

import (
	"io"

	"github.com/unrolled/render"
)

var renderer *render.Render

func init() {
	renderer = render.New(render.Options{
		Layout: "layout",
	})
}

// Render writes a rendered template to an io.Writer
func Render(w io.Writer, template string, data interface{}) error {
	if err := renderer.HTML(w, 0, template, data); err != nil {
		return err
	}

	return nil
}
