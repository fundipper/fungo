package plugin

import (
	"bytes"

	mermaid "github.com/abhinav/goldmark-mermaid"
	toc "github.com/abhinav/goldmark-toc"
	htmls "github.com/alecthomas/chroma/formatters/html"
	"github.com/fundipper/fungo/conf"
	images "github.com/fundipper/goldmark-images"
	links "github.com/fundipper/goldmark-links"
	videos "github.com/fundipper/goldmark-videos"
	mathjax "github.com/litao91/goldmark-mathjax"
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
	if !conf.PARSE_STATE {
		return
	}

	ext := []goldmark.Extender{}
	ext = append(ext, meta.Meta)

	mx := conf.NewConfig().Site.Markdown
	// state gfm
	if mx.State.GFM {
		ext = append(ext, extension.GFM)
	} else {
		if mx.State.Table {
			ext = append(ext, extension.Table)
		}
		if mx.State.Strikethrough {
			ext = append(ext, extension.Strikethrough)
		}
		if mx.State.Linkify {
			ext = append(ext, extension.Linkify)
		}
		if mx.State.TaskList {
			ext = append(ext, extension.TaskList)
		}
	}
	// state other
	if mx.State.Emoji {
		ext = append(ext, emoji.Emoji)
	}
	if mx.State.DefinitionList {
		ext = append(ext, extension.DefinitionList)
	}
	if mx.State.Footnote {
		ext = append(ext, extension.Footnote)
	}
	if mx.State.Typographer {
		ext = append(ext, extension.Typographer)
	}
	if mx.State.Mathjax {
		ext = append(ext, mathjax.MathJax)
	}
	if mx.State.Mermaid {
		ext = append(ext, &mermaid.Extender{})
	}

	// highlighting
	if mx.Highlighting.State {
		ext = append(ext, highlighting.NewHighlighting(
			highlighting.WithStyle(mx.Highlighting.Theme),
			highlighting.WithFormatOptions(
				htmls.WithLineNumbers(mx.Highlighting.LineNumber),
			),
		))
	}

	// images
	if mx.Image.State {
		ext = append(ext, images.NewExtender(
			mx.Image.Source,
			mx.Image.Target,
			mx.Image.Attribute,
		))
	}

	// links
	if mx.Link.State {
		ext = append(ext, links.NewExtender(
			mx.Link.Source,
			conf.NewConfig().Site.Markdown.Link.Attribute,
		))
	}

	// videos
	if mx.Video.State {
		ext = append(ext, videos.NewExtender(
			mx.Video.Source,
			mx.Video.Attribute,
		))
	}

	md = goldmark.New(
		goldmark.WithExtensions(ext...),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
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
