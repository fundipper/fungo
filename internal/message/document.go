package message

import (
	"net/http"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
	"github.com/fundipper/fungo/pkg/plugin"
	"github.com/julienschmidt/httprouter"
)

type Document struct {
	Model *conf.Model
}

func NewDocument(model *conf.Model) *Document {
	return &Document{
		Model: model,
	}
}

func (d *Document) Serve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result, err := compose.NewMarkdown().Item(r.RequestURI)
	if err != nil {
		panic(err)
	}
	err = plugin.NewHTML().Render(w, d.Model.Name, &Message{
		Path:    r.RequestURI,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Message: result,
		Sidebar: d.Model.Sidebar,
	})
	panic(err)
}

func (d *Document) Build(path string) error {
	result, err := compose.NewMarkdown().Item(path)
	if err != nil {
		panic(err)
	}
	return plugin.NewHTML().Export(path, d.Model.Name, &Message{
		Path:    path,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Message: result,
		Sidebar: d.Model.Sidebar,
	})
}
