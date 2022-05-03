package message

import (
	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
)

type Message struct {
	Lang    string
	Path    string
	Site    *conf.Site
	Theme   *conf.Theme
	Message *compose.Markdown
	Page    *compose.Page
	Catalog []*compose.Markdown
	Sidebar []*conf.Node
}
