package message

import (
	"net/http"
	"net/url"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
	"github.com/fundipper/fungo/internal/x/parse"
	"github.com/fundipper/fungo/pkg/cache"
	"github.com/fundipper/fungo/pkg/plugin"
	"github.com/julienschmidt/httprouter"
)

type Catalog struct{}

func NewCatalog() *Catalog {
	return &Catalog{}
}

func (c *Catalog) Serve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	path, err := url.QueryUnescape(r.RequestURI)
	if err != nil {
		panic(err)
	}

	key := parse.NewKey().Page(path)
	data, ok := cache.NewHash().Get(key)
	if !ok {
		panic(ok)
	}

	option := data.(*parse.Option)
	result, err := compose.NewMarkdown().List(option.Catalog, option.Page, conf.NewConfig().Site.Size)
	if err != nil {
		panic(err)
	}

	page, err := compose.NewCompute().Page(option.Catalog, option.Page)
	if err != nil {
		panic(err)
	}
	err = plugin.NewHTML().Render(w, option.Template, &Message{
		Path:    r.RequestURI,
		Lang:    option.Lang,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Catalog: result,
		Page:    page,
	})
	panic(err)
}

func (c *Catalog) Build(path string) error {
	key := parse.NewKey().Page(path)
	data, ok := cache.NewHash().Get(key)
	if !ok {
		panic(ok)
	}

	option := data.(*parse.Option)
	result, err := compose.NewMarkdown().List(option.Catalog, option.Page, conf.NewConfig().Site.Size)
	if err != nil {
		panic(err)
	}

	page, err := compose.NewCompute().Page(option.Catalog, option.Page)
	if err != nil {
		panic(err)
	}
	return plugin.NewHTML().Export(path, option.Template, &Message{
		Path:    path,
		Lang:    option.Lang,
		Site:    conf.NewConfig().Site,
		Theme:   conf.NewConfig().Theme,
		Catalog: result,
		Page:    page,
	})
}
