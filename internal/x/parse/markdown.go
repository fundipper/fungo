package parse

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/pkg/cache"
	"github.com/fundipper/fungo/pkg/plugin"
	"github.com/fundipper/fungo/pkg/util"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

type (
	Markdown struct {
		Model *conf.Model
	}

	Option struct {
		Template string
		Catalog  string
		Lang     string
		Page     int
	}
)

func NewMarkdown(model *conf.Model) *Markdown {
	return &Markdown{
		Model: model,
	}
}

func (m *Markdown) Parse(path string) error {
	source, err := util.NewTree().ReadFile(path)
	if err != nil {
		return err
	}

	ctx := parser.NewContext()
	content, err := plugin.NewMarkdown().Content(ctx, source)
	if err != nil {
		return err
	}

	toc, err := plugin.NewMarkdown().Table(ctx, source)
	if err != nil {
		return errors.New(conf.ERROR_TOC)
	}

	mx := meta.Get(ctx)
	if mx == nil {
		return errors.New(conf.ERROR_META)
	}

	// set route
	var lang string
	if mx[conf.META_LANG] != nil {
		lang = mx[conf.META_LANG].(string)
	}

	var slug string
	if mx[conf.META_SLUG] != nil {
		slug = mx[conf.META_SLUG].(string)
	}

	dir := strings.TrimPrefix(path, conf.PREFIX_CONTENT)
	if m.Model.Root {
		dir = strings.TrimPrefix(path, fmt.Sprintf(conf.PREFIX_PAGE, m.Model.Name))
	}
	route := NewPath().Route(lang, dir, slug)

	// set content
	_ = cache.NewString().Set(NewKey().Content(route), content.String())

	// set path
	_ = cache.NewString().Set(NewKey().Path(route), path)

	// set route
	_ = cache.NewSet().Set(m.Model.Name, route)

	// set meta
	_ = cache.NewHash().Set(route, mx)

	// set toc
	if toc != nil {
		_ = cache.NewString().Set(NewKey().TOC(route), toc.String())
	}

	// set lang
	if lang != "" {
		_ = cache.NewString().Set(NewKey().Lang(route), lang)
	}

	// check catalog
	if !m.Model.Catalog {
		return nil
	}

	var t time.Time
	// set archive
	if mx[conf.META_DATE] != nil {
		t, err = time.ParseInLocation("2006-01-02 15:04:05",
			fmt.Sprintf("%s 00:00:00", mx[conf.META_DATE].(string)),
			time.Local,
		)
		if err != nil {
			return err
		}
		_ = cache.NewString().Set(NewKey().Date(route), t.String())

		archive := NewPath().Archive(lang, fmt.Sprintf("%04d%02d", t.Year(), t.Month()))
		// set list
		total, ok := cache.NewZset().Set(archive, route, t.Unix())
		if ok {
			number := total / conf.NewConfig().Site.Size

			page := NewPath().Page(archive, number)
			// set route
			_ = cache.NewSet().Set(conf.META_ARCHIVE, page)
			// set option
			_ = cache.NewHash().Set(NewKey().Page(page), &Option{
				Template: conf.META_ARCHIVE,
				Catalog:  archive,
				Lang:     lang,
				Page:     number,
			})
		}
	}

	// set category
	if mx[conf.META_CATEGORY] != nil {
		category := NewPath().Category(lang, mx[conf.META_CATEGORY].(string))
		// set list
		total, ok := cache.NewZset().Set(category, route, t.Unix())
		if ok {
			number := total / conf.NewConfig().Site.Size

			page := NewPath().Page(category, number)
			// set route
			_ = cache.NewSet().Set(conf.META_CATEGORY, page)
			// set option
			_ = cache.NewHash().Set(NewKey().Page(page), &Option{
				Template: conf.META_CATEGORY,
				Catalog:  category,
				Lang:     lang,
				Page:     number,
			})
		}

	}

	// set tag
	if mx[conf.META_TAG] != nil {
		for _, v := range mx[conf.META_TAG].([]interface{}) {
			tag := NewPath().Tag(lang, v.(string))
			// set list
			total, ok := cache.NewZset().Set(tag, route, t.Unix())
			if ok {
				number := total / conf.NewConfig().Site.Size

				page := NewPath().Page(tag, number)
				// set route
				_ = cache.NewSet().Set(conf.META_TAG, page)
				// set option
				_ = cache.NewHash().Set(NewKey().Page(page), &Option{
					Template: conf.META_TAG,
					Catalog:  tag,
					Lang:     lang,
					Page:     number,
				})
			}
		}
	}

	catalog := NewPath().Catalog(lang, m.Model.Name)
	// set list
	total, ok := cache.NewZset().Set(catalog, route, t.Unix())
	if ok {
		number := total / conf.NewConfig().Site.Size

		page := NewPath().Page(catalog, number)
		// set catalog
		_ = cache.NewSet().Set(conf.META_CATALOG, page)
		// set option
		_ = cache.NewHash().Set(NewKey().Page(page), &Option{
			Template: m.Model.Template,
			Catalog:  catalog,
			Lang:     lang,
			Page:     number,
		})
	}
	return nil
}
