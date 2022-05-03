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

func (c *Customize) Serve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result, err := compose.NewMarkdown().List(c.Model.Action, 0, conf.NewConfig().Site.Size)
	if err != nil {
		panic(err)
	}
	err = plugin.NewHTML().Render(w, c.Model.Name, &Message{
		Path:    r.RequestURI,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Catalog: result,
	})
	panic(err)
}

func (c *Customize) Build(path string) error {
	result, err := compose.NewMarkdown().List(c.Model.Action, 0, conf.NewConfig().Site.Size)
	if err != nil {
		return err
	}
	return plugin.NewHTML().Export(path, c.Model.Name, &Message{
		Path:    path,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Catalog: result,
	})
}
