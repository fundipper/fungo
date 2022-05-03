package compose

import (
	"errors"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/parse"
	"github.com/fundipper/fungo/pkg/cache"
	"github.com/fundipper/fungo/pkg/util"
)

type Compute struct{}

func NewCompute() *Compute {
	return &Compute{}
}

func (c *Compute) Page(model string, index int) (*Page, error) {
	total, ok := cache.NewZset().Len(model)
	if !ok {
		return nil, errors.New(conf.ERROR_RANK)
	}

	total = total / conf.NewConfig().Site.Size
	factor := conf.NewConfig().Site.Size / 2

	pre := 0
	next := 0
	if index-1 >= 0 {
		pre = index - 1
	}

	if index+1 <= total {
		next = index + 1
	}

	start := index - factor
	end := index + factor
	if start < 0 {
		start = 0
		if total-conf.NewSite().Size >= 0 {
			end = conf.NewConfig().Site.Size
		}
	}
	if end > total {
		end = total
		if total-conf.NewConfig().Site.Size >= 0 {
			start = total - conf.NewConfig().Site.Size
		}
	}

	data := []int{}
	for i := start; i < end; i++ {
		data = append(data, i)
	}
	return &Page{
		Index: index,
		Size:  conf.NewConfig().Site.Size,
		Total: total,
		Pre:   pre,
		Next:  next,
		Data:  data,
		Path:  model,
	}, nil
}

func (c *Compute) Top(path string) ([]*Markdown, error) {
	model := util.NewPath().Dir(path)
	index, ok := cache.NewZset().Rank(model, path)
	if !ok {
		return nil, errors.New(conf.ERROR_RANK)
	}

	start := index + 1
	end := index + conf.NewConfig().Site.Amount
	data, ok := cache.NewZset().Get(model, start, end)
	if !ok {
		return nil, errors.New(conf.ERROR_CATALOG)
	}

	result := []*Markdown{}
	for _, v := range data {
		item, ok := cache.NewHash().Get(v)
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
			Meta: item.(map[string]interface{}),
			Date: date,
			Lang: &lang,
		})
	}
	return result, nil
}
