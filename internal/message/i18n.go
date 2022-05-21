package message

import (
	"net/http"
	"path/filepath"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
	"github.com/fundipper/fungo/pkg/plugin"
	"github.com/julienschmidt/httprouter"
)

type I18N struct {
	Model *conf.Model
}

func NewI18N(model *conf.Model) *I18N {
	return &I18N{
		Model: model,
	}
}

func (i *I18N) Serve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result, err := compose.NewMarkdown().List(i.Model.Action, 0, conf.NewConfig().Site.Size)
	if err != nil {
		panic(err)
	}
	err = plugin.NewHTML().Render(w, i.Model.Name, &Message{
		Path:    r.RequestURI,
		Lang:    filepath.Base(r.RequestURI),
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Catalog: result,
	})
	panic(err)
}

func (i *I18N) Build(path string) error {
	result, err := compose.NewMarkdown().List(i.Model.Action, 0, conf.NewConfig().Site.Size)
	if err != nil {
		return err
	}
	return plugin.NewHTML().Export(path, i.Model.Name, &Message{
		Path:    path,
		Lang:    filepath.Base(path),
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Catalog: result,
	})
}
