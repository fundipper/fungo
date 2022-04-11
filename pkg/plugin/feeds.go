package plugin

import (
	"encoding/json"

	"github.com/beevik/etree"
)

type (
	Feeds struct {
		Title    string   `json:"title,omitempty"`
		Link     string   `json:"url,omitempty"`
		ID       string   `json:"id,omitempty"`
		Summary  string   `json:"summary,omitempty"`
		Content  string   `json:"content,omitempty"`
		Created  string   `json:"date_published,omitempty"`
		Language string   `json:"language,omitempty"`
		Author   *Author  `json:"author,omitempty"`
		Entry    []*Feeds `json:"items,omitempty"`
	}
	Author struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	}
)

func NewFeeds() *Feeds {
	return &Feeds{}
}

func (f *Feeds) Atom(args *Feeds) (string, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	feed := doc.CreateElement("feed")
	feed.CreateAttr("xmlns", "http://www.w3.org/2005/Atom")

	if args.Title != "" {
		title := feed.CreateElement("title")
		title.CreateText(args.Title)
	}

	if args.Link != "" {
		link := feed.CreateElement("link")
		link.CreateAttr("href", args.Link)
	}

	if args.Summary != "" {
		summary := feed.CreateElement("summary")
		summary.CreateAttr("type", "html")
		summary.CreateText(args.Summary)
	}

	if args.Created != "" {
		updated := feed.CreateElement("updated")
		updated.CreateText(args.Created)
	}

	for _, v := range args.Entry {
		entry := feed.CreateElement("entry")

		if v.Title != "" {
			title := entry.CreateElement("title")
			title.CreateText(v.Title)
		}

		if v.Link != "" {
			link := entry.CreateElement("link")
			link.CreateAttr("href", v.Link)
		}

		if v.ID != "" {
			id := entry.CreateElement("id")
			id.CreateText(v.Link)
		}

		if v.Summary != "" {
			summary := entry.CreateElement("summary")
			summary.CreateAttr("type", "html")
			summary.CreateText(v.Summary)
		}

		if v.Content != "" {
			content := entry.CreateElement("content")
			content.CreateAttr("type", "html")
			content.CreateText(v.Content)
		}

		if v.Created != "" {
			updated := entry.CreateElement("updated")
			updated.CreateText(v.Created)
		}

		if v.Language != "" {
			language := entry.CreateElement("language")
			language.CreateText(v.Language)
		}

		if v.Author != nil {
			author := entry.CreateElement("author")

			if v.Author.Name != "" {
				name := author.CreateElement("name")
				name.CreateText(v.Author.Name)
			}

			if v.Author.Email != "" {
				email := author.CreateElement("email")
				email.CreateText(v.Author.Email)
			}
		}
	}

	doc.Indent(4)
	return doc.WriteToString()
}

func (f *Feeds) RSS(args *Feeds) (string, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	rss := doc.CreateElement("rss")
	rss.CreateAttr("version", "2.0")

	channel := rss.CreateElement("channel")

	if args.Title != "" {
		title := channel.CreateElement("title")
		title.CreateText(args.Title)
	}

	if args.Link != "" {
		link := channel.CreateElement("link")
		link.CreateAttr("href", args.Link)
	}

	if args.Summary != "" {
		description := channel.CreateElement("description")
		description.CreateText(args.Summary)
	}

	if args.Created != "" {
		pubDate := channel.CreateElement("pubDate")
		pubDate.CreateText(args.Created)
	}

	for _, v := range args.Entry {
		item := channel.CreateElement("item")

		if v.Title != "" {
			title := item.CreateElement("title")
			title.CreateText(v.Title)
		}

		if v.Link != "" {
			link := item.CreateElement("link")
			link.CreateAttr("href", v.Link)
		}

		if v.Summary != "" {
			description := item.CreateElement("description")
			description.CreateText(v.Summary)
		}

		if v.Content != "" {
			content := item.CreateElement("content")
			content.CreateText(v.Content)
		}

		if v.Created != "" {
			pubDate := item.CreateElement("pubDate")
			pubDate.CreateText(v.Created)
		}

		if v.Language != "" {
			language := item.CreateElement("language")
			language.CreateText(v.Language)
		}
	}

	doc.Indent(4)
	return doc.WriteToString()
}

func (f *Feeds) JSON(args *Feeds) (string, error) {
	data := map[string]interface{}{
		"version":       "https://jsonfeed.org/version/1",
		"title":         args.Title,
		"home_page_url": args.Link,
		"description":   args.Summary,
		"items":         args.Entry,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
