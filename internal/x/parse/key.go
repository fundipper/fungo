package parse

import (
	"fmt"

	"github.com/fundipper/fungo/conf"
)

type Key struct{}

func NewKey() *Key {
	return &Key{}
}

func (k *Key) Content(route string) string {
	return fmt.Sprintf("%s:%s", conf.MESSAGE_CONTENT, route)
}

func (k *Key) TOC(route string) string {
	return fmt.Sprintf("%s:%s", conf.MESSAGE_TOC, route)
}

func (k *Key) Date(route string) string {
	return fmt.Sprintf("%s:%s", conf.META_DATE, route)
}

func (k *Key) Lang(route string) string {
	return fmt.Sprintf("%s:%s", conf.META_LANG, route)
}
