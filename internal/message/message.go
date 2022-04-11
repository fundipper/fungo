package message

import (
	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
)

type Message struct {
	Path    string
	Site    *conf.Site
	Theme   *conf.Theme
	Message *compose.Markdown
	Catalog *compose.Catalog
	Sidebar []*conf.Node
}
