package cli

import (
	"fmt"
	"sync"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/message"
	"github.com/fundipper/fungo/pkg/cache"
	"github.com/otiai10/copy"
)

type Build struct{}

func NewBuild() *Build {
	once.Do(func() {
		NewPrase().Run()
	})
	return &Build{}
}

func (b *Build) Run() {
	var wg sync.WaitGroup

	for _, v := range conf.NewConfig().Article {
		route, ok := cache.NewSet().Get(v.Name)
		if !ok {
			continue
		}
		for _, item := range route {
			wg.Add(1)
			go func(model *conf.Model, path string) {
				defer wg.Done()

				_ = message.NewArticle(model).Build(path)
			}(v, item)
		}
	}

	for _, v := range conf.NewConfig().Collection {
		route, ok := cache.NewSet().Get(v.Name)
		if !ok {
			continue
		}
		for _, item := range route {
			wg.Add(1)
			go func(model *conf.Model, path string) {
				defer wg.Done()

				_ = message.NewCollection(model).Build(path)
			}(v, item)
		}

		if v.Contents {
			route, ok := cache.NewSet().Get(conf.META_CONTENTS)
			if !ok {
				continue
			}

			for _, item := range route {
				wg.Add(1)
				go func(model *conf.Model, path string) {
					defer wg.Done()

					_ = message.NewContents(model).Build(path)
				}(v, item)
			}
		}
	}

	for _, v := range conf.NewConfig().Document {
		route, ok := cache.NewSet().Get(v.Name)
		if !ok {
			continue
		}
		for _, item := range route {
			wg.Add(1)
			go func(model *conf.Model, path string) {
				defer wg.Done()

				_ = message.NewDocument(model).Build(path)
			}(v, item)
		}
	}

	for _, v := range conf.NewConfig().Page {
		route, ok := cache.NewSet().Get(v.Name)
		if !ok {
			continue
		}
		for _, item := range route {
			wg.Add(1)
			go func(model *conf.Model, path string) {
				defer wg.Done()

				_ = message.NewPage(model).Build(path)
			}(v, item)
		}
	}

	for _, v := range []string{conf.META_ARCHIVE, conf.META_CATEGORY, conf.META_TAG, conf.META_CATALOG} {
		route, ok := cache.NewSet().Get(v)
		if !ok {
			continue
		}
		for _, item := range route {
			wg.Add(1)
			go func(path string) {
				defer wg.Done()

				_ = message.NewCatalog().Build(path)
			}(item)
		}
	}

	for _, v := range conf.NewConfig().Customize {
		wg.Add(1)
		go func(model *conf.Model, path string) {
			defer wg.Done()

			_ = message.NewCustomize(model).Build(path)
		}(v, v.Route)
	}

	for _, v := range conf.NewConfig().I18N {
		wg.Add(1)
		go func(model *conf.Model, path string) {
			defer wg.Done()

			_ = message.NewI18N(model).Build(path)
		}(v, v.Route)
	}

	_, _ = message.NewSitemap().BuildList()
	for _, item := range []string{conf.MODEL_ARTICLE, conf.MODEL_DOCUMENT, conf.MODEL_PAGE, conf.MODEL_CUSTOMIZE, conf.MODEL_I18N, conf.META_CATALOG} {
		_, _ = message.NewSitemap().Build(item)
	}

	for _, item := range []string{conf.FEEDS_ATOM, conf.FEEDS_RSS, conf.FEEDS_JSON} {
		_, _ = message.NewFeeds(conf.NewConfig().Site.Feeds.Action).Build(item)
	}

	_ = copy.Copy(conf.CONTENT_MEDIA, conf.PUBLIC_MEIDA)
	_ = copy.Copy(conf.THEME_ASSETS, conf.PUBLIC_ASSETS)

	fmt.Println("build file successed")
	wg.Wait()
}
