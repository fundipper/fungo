package compose

import (
	"fmt"
	"time"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/pkg/cache"
	"github.com/fundipper/fungo/pkg/plugin"
)

type Sitemap struct {
	Now    string
	Origin string
}

func NewSitemap() *Sitemap {
	return &Sitemap{
		Now:    time.Now().Format("2006-01-02"),
		Origin: conf.NewConfig().Site.Origin,
	}
}

func (s *Sitemap) List() (string, error) {
	result := []*plugin.Sitemap{}
	for _, v := range []string{conf.MODEL_ARTICLE, conf.MODEL_DOCUMENT, conf.MODEL_PAGE, conf.MODEL_I18N, conf.MODEL_CUSTOMIZE, conf.META_CATALOG} {
		result = append(result, &plugin.Sitemap{
			Loc:     fmt.Sprintf(conf.SITEMAP_XML, s.Origin, v),
			Lastmod: s.Now,
		})
	}

	return plugin.NewSitemap().List(result)
}

func (s *Sitemap) Item(name string) (string, error) {
	result := []*plugin.Sitemap{}
	switch name {
	case conf.MODEL_ARTICLE:
		for _, v := range conf.NewConfig().Article {
			route, ok := cache.NewSet().Get(v.Name)
			if !ok {
				continue
			}
			for _, item := range route {
				result = append(result, &plugin.Sitemap{
					Loc:        fmt.Sprintf("%s%s", s.Origin, item),
					Lastmod:    s.Now,
					Changefreq: conf.NewConfig().Site.Sitemap.Changefreq,
					Priority:   conf.NewConfig().Site.Sitemap.Priority,
				})
			}
		}
	case conf.MODEL_COLLECTION:
		for _, v := range conf.NewConfig().Collection {
			route, ok := cache.NewSet().Get(v.Name)
			if !ok {
				continue
			}
			for _, item := range route {
				result = append(result, &plugin.Sitemap{
					Loc:        fmt.Sprintf("%s%s", s.Origin, item),
					Lastmod:    s.Now,
					Changefreq: conf.NewConfig().Site.Sitemap.Changefreq,
					Priority:   conf.NewConfig().Site.Sitemap.Priority,
				})
			}
		}
	case conf.MODEL_DOCUMENT:
		for _, v := range conf.NewConfig().Document {
			route, ok := cache.NewSet().Get(v.Name)
			if !ok {
				continue
			}
			for _, item := range route {
				result = append(result, &plugin.Sitemap{
					Loc:        fmt.Sprintf("%s%s", s.Origin, item),
					Lastmod:    s.Now,
					Changefreq: conf.NewConfig().Site.Sitemap.Changefreq,
					Priority:   conf.NewConfig().Site.Sitemap.Priority,
				})
			}
		}
	case conf.MODEL_PAGE:
		for _, v := range conf.NewConfig().Page {
			route, ok := cache.NewSet().Get(v.Name)
			if !ok {
				continue
			}
			for _, item := range route {
				result = append(result, &plugin.Sitemap{
					Loc:        fmt.Sprintf("%s%s", s.Origin, item),
					Lastmod:    s.Now,
					Changefreq: conf.NewConfig().Site.Sitemap.Changefreq,
					Priority:   conf.NewConfig().Site.Sitemap.Priority,
				})
			}
		}
	case conf.MODEL_CUSTOMIZE:
		for _, v := range conf.NewConfig().Customize {
			result = append(result, &plugin.Sitemap{
				Loc:        fmt.Sprintf("%s%s", s.Origin, v.Route),
				Lastmod:    s.Now,
				Changefreq: conf.NewConfig().Site.Sitemap.Changefreq,
				Priority:   conf.NewConfig().Site.Sitemap.Priority,
			})
		}
	case conf.MODEL_I18N:
		for _, v := range conf.NewConfig().I18N {
			result = append(result, &plugin.Sitemap{
				Loc:        fmt.Sprintf("%s%s", s.Origin, v.Route),
				Lastmod:    s.Now,
				Changefreq: conf.NewConfig().Site.Sitemap.Changefreq,
				Priority:   conf.NewConfig().Site.Sitemap.Priority,
			})
		}
	case conf.META_CATALOG:
		for _, v := range []string{conf.META_ARCHIVE, conf.META_CATEGORY, conf.META_TAG, conf.META_CATALOG} {
			route, ok := cache.NewSet().Get(v)
			if !ok {
				continue
			}
			for _, item := range route {
				result = append(result, &plugin.Sitemap{
					Loc:        fmt.Sprintf("%s%s", s.Origin, item),
					Lastmod:    s.Now,
					Changefreq: conf.NewConfig().Site.Sitemap.Changefreq,
					Priority:   conf.NewConfig().Site.Sitemap.Priority,
				})
			}
		}
	}
	return plugin.NewSitemap().Item(result)
}
