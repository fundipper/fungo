package plugin

import (
	"path/filepath"

	"github.com/fundipper/fungo/conf"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func init() {
	if !conf.PARSE_STATE {
		return
	}

	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc(conf.CONFIG_TYPE, toml.Unmarshal)

	for _, v := range conf.NewConfig().I18N {
		path := filepath.Join(conf.THEME_I18N, v.Path)
		NewI18N().Bundle.MustLoadMessageFile(path)
	}
}

type I18N struct {
	Bundle *i18n.Bundle
}

func NewI18N() *I18N {
	return &I18N{
		Bundle: bundle,
	}
}

func (i *I18N) Parse(lang, s string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: s,
	})
}
