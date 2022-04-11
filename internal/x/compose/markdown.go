package compose

import (
	"sort"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/parse"
	"github.com/fundipper/fungo/pkg/cache"
)

type (
	Markdown struct {
		Name    string
		Content string
		Date    string
		Meta    map[string]interface{}
		TOC     *TOC
		Lang    *Lang
	}

	TOC struct {
		State   bool
		Content string
	}

	Lang struct {
		State   bool
		Content string
	}
)

func NewMarkdown() *Markdown {
	return &Markdown{}
}

func (m *Markdown) List(model string) *Catalog {
	data, ok := cache.NewList().Get(model)
	if !ok {
		panic(conf.ERROR_CATALOG)
	}

	result := Catalog{}
	for _, v := range data {
		meta, ok := cache.NewHash().Get(v)
		if !ok {
			continue
		}

		key := parse.NewKey().Date(v)
		date, ok := cache.NewString().Get(key)
		if !ok {
			continue
		}

		key = parse.NewKey().Lang(v)
		lang, ok := cache.NewString().Get(key)
		result = append(result, &Markdown{
			Name: v,
			Meta: meta,
			Date: date,
			Lang: &Lang{
				State:   ok,
				Content: lang,
			},
		})
	}
	sort.Sort(result)
	return &result
}

func (m *Markdown) Item(path string) *Markdown {

	meta, ok := cache.NewHash().Get(path)
	if !ok {
		panic(conf.ERROR_META)
	}

	key := parse.NewKey().Content(path)
	content, ok := cache.NewString().Get(key)
	if !ok {
		panic(conf.ERROR_CONTENT)
	}

	key = parse.NewKey().TOC(path)
	toc, ok1 := cache.NewString().Get(key)

	key = parse.NewKey().Lang(path)
	lang, ok2 := cache.NewString().Get(key)
	return &Markdown{
		Meta:    meta,
		Content: content,
		TOC: &TOC{
			State:   ok1,
			Content: toc,
		},
		Lang: &Lang{
			State:   ok2,
			Content: lang,
		},
	}
}
