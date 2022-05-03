package compose

import (
	"errors"
	"net/url"

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
		Page    *Page
		Top     []*Markdown
	}

	TOC struct {
		State   bool
		Content string
	}

	Lang struct {
		State   bool
		Content string
	}

	Page struct {
		Index int
		Size  int
		Total int
		Pre   int
		Next  int
		Data  []int
		Path  string
	}
)

func NewMarkdown() *Markdown {
	return &Markdown{}
}

func (m *Markdown) List(model string, page, size int) ([]*Markdown, error) {
	start := page*size + 1
	end := (page + 1) * size
	data, ok := cache.NewZset().Get(model, start, end)
	if !ok {
		return nil, errors.New(conf.ERROR_CATALOG)
	}

	result := []*Markdown{}
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

		lang := Lang{}
		key = parse.NewKey().Lang(v)
		if content, ok := cache.NewString().Get(key); ok {
			lang.Content = content
			lang.State = ok
		}

		result = append(result, &Markdown{
			Name: v,
			Meta: meta.(map[string]interface{}),
			Date: date,
			Lang: &lang,
		})
	}
	return result, nil
}

func (m *Markdown) Item(path string) (*Markdown, error) {
	path, err := url.QueryUnescape(path)
	if err != nil {
		return nil, err
	}

	meta, ok := cache.NewHash().Get(path)
	if !ok {
		return nil, errors.New(conf.ERROR_META)
	}

	key := parse.NewKey().Content(path)
	content, ok := cache.NewString().Get(key)
	if !ok {
		return nil, errors.New(conf.ERROR_CONTENT)
	}

	toc := TOC{}
	key = parse.NewKey().TOC(path)
	if content, ok := cache.NewString().Get(key); ok {
		toc.Content = content
		toc.State = ok
	}

	lang := Lang{}
	key = parse.NewKey().Lang(path)
	if content, ok := cache.NewString().Get(key); ok {
		lang.Content = content
		lang.State = ok
	}
	return &Markdown{
		Meta:    meta.(map[string]interface{}),
		Content: content,
		TOC:     &toc,
		Lang:    &lang,
	}, nil
}
