package cli

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/message"
	"github.com/fundipper/fungo/pkg/cache"
	"github.com/julienschmidt/httprouter"
)

type Serve struct{}

func NewServe() *Serve {
	once.Do(func() {
		NewPrase().Run()
	})
	return &Serve{}
}

func (s *Serve) Run() {
	r := httprouter.New()

	for _, v := range conf.NewConfig().Article {
		route, ok := cache.NewSet().Get(v.Name)
		if !ok {
			continue
		}
		for item := range route {
			r.GET(item, message.NewArticle(v).Serve)
		}
	}

	for _, v := range conf.NewConfig().Document {
		route, ok := cache.NewSet().Get(v.Name)
		if !ok {
			continue
		}
		for item := range route {
			r.GET(item, message.NewDocument(v).Serve)
		}
	}

	for _, v := range conf.NewConfig().Page {
		route, ok := cache.NewSet().Get(v.Name)
		if !ok {
			continue
		}
		for item := range route {
			r.GET(item, message.NewPage(v).Serve)
		}
	}

	for _, v := range []string{conf.META_ARCHIVE, conf.META_CATEGORY, conf.META_TAG, conf.META_CATALOG} {
		route, ok := cache.NewSet().Get(v)
		if !ok {
			continue
		}
		for item := range route {
			r.GET(item, message.NewCatalog(v).Serve)
		}
	}

	for _, v := range conf.NewConfig().Customize {
		r.GET(v.Route, message.NewCustomize(v).Serve)
	}

	for _, v := range conf.NewConfig().I18N {
		r.GET(v.Route, message.NewI18N(v).Serve)
	}

	r.GET("/sitemap.xml", message.NewSitemap().ServeList)
	r.GET("/sitemap/:name", message.NewSitemap().Serve)
	r.GET("/atom.xml", message.NewFeeds(conf.NewConfig().Site.Feeds.Action).Serve)
	r.GET("/rss.xml", message.NewFeeds(conf.NewConfig().Site.Feeds.Action).Serve)
	r.GET("/feeds.json", message.NewFeeds(conf.NewConfig().Site.Feeds.Action).Serve)
	r.ServeFiles("/media/*filepath", http.Dir(conf.CONTENT_MEDIA))
	r.ServeFiles("/assets/*filepath", http.Dir(conf.THEME_ASSETS))

	r.NotFound = http.FileServer(http.Dir(conf.PUBLIC_ROOT))
	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprintf(w, "error: %v", i)
	}

	color.Black(_SERVER, conf.NewConfig().Site.Port)
	log.Fatal(http.ListenAndServe(conf.NewConfig().Site.Port, r))
}
