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
	result, err := NewMarkdown().List(f.Model, 0, conf.NewConfig().Site.Feeds.Limit)
	if err != nil {
		return "", err
	}

	message := &plugin.Feeds{
		Title:   conf.NewConfig().Site.Name,
		Link:    conf.NewConfig().Site.Origin,
		Summary: conf.NewConfig().Site.Slogan,
		Created: time.Now().String(),
	}

	data := []*plugin.Feeds{}
	for _, v := range result {
		var title string
		if v.Meta[conf.META_TITLE] != nil {
			title = v.Meta[conf.META_TITLE].(string)
		}

		var description string
		if v.Meta[conf.META_DESCRIPTION] != nil {
			description = v.Meta[conf.META_DESCRIPTION].(string)
		}

		var author string
		if v.Meta[conf.META_AUTHOR] != nil {
			author = v.Meta[conf.META_AUTHOR].(string)
		}
		item := &plugin.Feeds{
			Title:   title,
			Link:    fmt.Sprintf("%s%s", conf.NewConfig().Site.Origin, v.Name),
			Summary: description,
			Created: v.Date,
			Author: &plugin.Author{
				Name: author,
			},
		}

		if conf.NewConfig().Site.Feeds.Content {
			message, err := NewMarkdown().Item(v.Name)
			if err != nil {
				continue
			}

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
