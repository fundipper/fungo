package message

import (
	"net/http"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
	"github.com/fundipper/fungo/pkg/plugin"
	"github.com/julienschmidt/httprouter"
)

type Article struct {
	Model *conf.Model
}

func NewArticle(model *conf.Model) *Article {
	return &Article{
		Model: model,
	}
}

func (a *Article) Serve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := compose.NewMarkdown().Item(r.RequestURI)
	err := plugin.NewHTML().Render(w, a.Model.Name, &Message{
		Path:    r.RequestURI,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Message: result,
	})
	panic(err)
}

func (a *Article) Build(path string) error {
	result := compose.NewMarkdown().Item(path)
	return plugin.NewHTML().Export(path, a.Model.Name, &Message{
		Path:    path,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Message: result,
	})
}
