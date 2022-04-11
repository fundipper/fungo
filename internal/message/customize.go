package message

import (
	"net/http"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
	"github.com/fundipper/fungo/pkg/plugin"
	"github.com/julienschmidt/httprouter"
)

type Customize struct {
	Model *conf.Model
}

func NewCustomize(model *conf.Model) *Customize {
	return &Customize{
		Model: model,
	}
}

func (c *Customize) Serve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := compose.NewMarkdown().List(c.Model.Action)
	err := plugin.NewHTML().Render(w, c.Model.Name, &Message{
		Path:    r.RequestURI,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Catalog: result,
	})
	panic(err)
}

func (c *Customize) Build(path string) error {
	result := compose.NewMarkdown().List(c.Model.Action)
	return plugin.NewHTML().Export(path, c.Model.Name, &Message{
		Path:    path,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Catalog: result,
	})
}
