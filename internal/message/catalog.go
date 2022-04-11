package message

import (
	"net/http"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
	"github.com/fundipper/fungo/pkg/plugin"
	"github.com/julienschmidt/httprouter"
)

type Catalog struct {
	Model string
}

func NewCatalog(model string) *Catalog {
	return &Catalog{
		Model: model,
	}
}

func (c *Catalog) Serve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := compose.NewMarkdown().List(r.RequestURI)
	err := plugin.NewHTML().Render(w, c.Model, &Message{
		Path:    r.RequestURI,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Catalog: result,
	})
	panic(err)
}

func (c *Catalog) Build(path string) error {
	result := compose.NewMarkdown().List(path)
	return plugin.NewHTML().Export(path, c.Model, &Message{
		Path:    path,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Catalog: result,
	})
}
