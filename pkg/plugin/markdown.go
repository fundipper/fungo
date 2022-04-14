package plugin

import (
	"bytes"

	toc "github.com/abhinav/goldmark-toc"
	"github.com/fundipper/fungo/conf"
	images "github.com/fundipper/goldmark-images"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
)

var md goldmark.Markdown

func init() {
	if conf.PARSE_STATE {
		md = goldmark.New(
			goldmark.WithExtensions(
				extension.GFM,
				emoji.Emoji,
				meta.Meta,
				images.NewExtender(func(src string) (string, map[string]string) {
					return conf.NewSite().Markdown.Lazyload.Data.Value,
						map[string]string{
							conf.NewSite().Markdown.Lazyload.Class.Key: conf.NewSite().Markdown.Lazyload.Class.Value,
							conf.NewSite().Markdown.Lazyload.Data.Key:  src,
						}
				}),
				highlighting.NewHighlighting(
					highlighting.WithStyle(conf.NewConfig().Site.Markdown.Highlighting),
				),
			),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithXHTML(),
			),
		)
	}
}

type Markdown struct{}

func NewMarkdown() *Markdown {
	return &Markdown{}
}

func (m *Markdown) Content(ctx parser.Context, source []byte) (*bytes.Buffer, error) {
	result := bytes.Buffer{}
	err := md.Convert(source, &result, parser.WithContext(ctx))
	return &result, err
}

func (m *Markdown) Table(ctx parser.Context, source []byte) (*bytes.Buffer, error) {
	doc := md.Parser().Parse(text.NewReader(source))
	tree, err := toc.Inspect(doc, source)
	if err != nil {
		return nil, err
	}

	list := toc.RenderList(tree)
	if list == nil {
		return nil, nil
	}

	result := bytes.Buffer{}
	err = md.Renderer().Render(&result, source, list)
	return &result, err
}
