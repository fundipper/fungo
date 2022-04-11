package compose

import (
	"errors"
	"fmt"
	"time"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/pkg/plugin"
)

type Feeds struct {
	Model string
}

func NewFeeds(model string) *Feeds {
	return &Feeds{
		Model: model,
	}
}

func (f *Feeds) Generate(name string) (string, error) {
	result := NewMarkdown().List(f.Model)

	message := &plugin.Feeds{
		Title:   conf.NewConfig().Site.Name,
		Link:    conf.NewConfig().Site.Origin,
		Summary: conf.NewConfig().Site.Slogan,
		Created: time.Now().String(),
	}

	data := []*plugin.Feeds{}
	for _, v := range *result {
		item := &plugin.Feeds{
			Title:   v.Meta[conf.META_TITLE].(string),
			Link:    fmt.Sprintf("%s%s", conf.NewConfig().Site.Origin, v.Name),
			Summary: v.Meta[conf.META_DESCRIPTION].(string),
			Created: v.Date,
			Author: &plugin.Author{
				Name: v.Meta[conf.META_AUTHOR].(string),
			},
		}

		if conf.NewConfig().Site.Feeds.Content {
			message := NewMarkdown().Item(v.Name)

			item.Content = message.Content
			item.Language = message.Lang.Content
		}

		data = append(data, item)
	}

	message.Entry = data

	switch name {
	case conf.FEEDS_ATOM:
		return plugin.NewFeeds().Atom(message)
	case conf.FEEDS_RSS:
		return plugin.NewFeeds().RSS(message)
	case conf.FEEDS_JSON:
		return plugin.NewFeeds().JSON(message)
	default:
		return "", errors.New(conf.ERROR_MATCH)
	}
}
