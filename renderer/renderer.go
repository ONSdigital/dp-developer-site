package renderer

import (
	"html/template"
	"io"

	"github.com/ONSdigital/dp-developer-site/utils"
	"github.com/unrolled/render"
)

var renderer *render.Render

var funcMap = template.FuncMap{
	"hasEnums": utils.HasEnums,
	"join":     utils.Join,
}

func init() {
	renderer = render.New(render.Options{
		Layout: "layout",
		Funcs: []template.FuncMap{
			funcMap,
		},
	})
}

// Render writes a rendered template to an io.Writer
func Render(w io.Writer, template string, data interface{}) error {
	if err := renderer.HTML(w, 0, template, data); err != nil {
		return err
	}

	return nil
}
