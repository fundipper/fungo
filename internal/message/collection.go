package message

import (
	"net/http"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
	"github.com/fundipper/fungo/pkg/plugin"
	"github.com/julienschmidt/httprouter"
)

type Collection struct {
	Model *conf.Model
}

func NewCollection(model *conf.Model) *Collection {
	return &Collection{
		Model: model,
	}
}

func (c *Collection) Serve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result, err := compose.NewMarkdown().Item(r.RequestURI)
	if err != nil {
		panic(err)
	}

	err = plugin.NewHTML().Render(w, c.Model.Name, &Message{
		Path:    r.RequestURI,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Message: result,
	})
	panic(err)
}

func (c *Collection) Build(path string) error {
	result, err := compose.NewMarkdown().Item(path)
	if err != nil {
		panic(err)
	}

	return plugin.NewHTML().Export(path, c.Model.Name, &Message{
		Path:    path,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Message: result,
	})
}
