package conf

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const (
	CONFIG_ROOT  = "config"
	CONFIG_THEME = "theme"
	CONFIG_TYPE  = "toml"

	MESSAGE_CONTENT  = "content"
	MESSAGE_TOC      = "toc"
	META_CATALOG     = "catalog"
	META_ARCHIVE     = "archive"
	META_CATEGORY    = "category"
	META_DATE        = "date"
	META_TAG         = "tag"
	META_LANG        = "lang"
	META_SLUG        = "slug"
	META_TITLE       = "title"
	META_DESCRIPTION = "description"
	META_AUTHOR      = "author"

	MODEL_ARTICLE   = "article"
	MODEL_DOCUMENT  = "document"
	MODEL_PAGE      = "page"
	MODEL_I18N      = "i18n"
	MODEL_CUSTOMIZE = "customize"
	MODEL_STATIC    = "static"
	MODEL_SITE      = "site"
	MODEL_THMEM     = "theme"
	MODEL_EXTEND    = "extend"

	FEEDS_ATOM = "atom.xml"
	FEEDS_RSS  = "rss.xml"
	FEEDS_JSON = "feeds.json"

	CONTENT_ROOT   = "content"
	CONTENT_MEDIA  = "content/media"
	PUBLIC_MEIDA   = "public/media"
	PUBLIC_ASSETS  = "public/assets"
	PUBLIC_SITEMAP = "public/sitemap.xml"
	THEME_ROOT     = "theme"
	THEME_DEFAULT  = "fungo"

	ROOT_TOML    = "config.toml"
	CONTENT_MD   = "content/%s.md"
	PUBLIC_FEEDS = "public/%s"
	PUBLIC_HTML  = "public%sindex.html"
	PUBLIC_XML   = "public/sitemap/%s.xml"
	SOURCE_YAML  = "source/%s.yaml"
	SITEMAP_XML  = "%s/sitemap/%s.xml"

	PREFIX_CONTENT = "content/"
	PREFIX_PAGE    = "%s/"
	SUFFIX_MD      = ".md"

	URL_SITE  = "https://github.com/fundipper/site"
	URL_THEME = "https://github.com/fundipper/theme"

	ERROR_CONFIG  = "config is nil"
	ERROR_CONTENT = "content is nil"
	ERROR_META    = "meta is nil"
	ERROR_TOC     = "toc is nil"
	ERROR_CATALOG = "catalog is nil"
	ERROR_EXIST   = "file is exist"
	ERROR_MATCH   = "name not match"
)

var (
	v               *viper.Viper
	config          *Config
	PARSE_STATE     bool
	THEME_USED      string
	THEME_I18N      string
	THEME_ASSETS    string
	THEME_TEMPLATES string
	THEME_HTML      string
	THEME_TOML      string
)

func init() {
	v = viper.New()
	v.SetConfigName(CONFIG_ROOT)
	v.SetConfigType(CONFIG_TYPE)
	v.AddConfigPath("./")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	config = &Config{
		Article:   NewModel().Article(),
		Customize: NewModel().Customize(),
		Document:  NewModel().Document(),
		Page:      NewModel().Page(),
		I18N:      NewModel().I18N(),
		Static:    NewModel().Static(),
		Site:      NewSite(),
	}

	THEME_USED = fmt.Sprintf("theme/%s", config.Site.Theme)
	THEME_I18N = fmt.Sprintf("theme/%s/i18n", config.Site.Theme)
	THEME_ASSETS = fmt.Sprintf("theme/%s/assets", config.Site.Theme)
	THEME_TEMPLATES = fmt.Sprintf("theme/%s/templates", config.Site.Theme)
	THEME_HTML = fmt.Sprintf("theme/%s/templates/*.html", config.Site.Theme)
	THEME_TOML = fmt.Sprintf("theme/%s/theme.toml", config.Site.Theme)

	v.SetConfigName(CONFIG_THEME)
	v.SetConfigType(CONFIG_TYPE)
	v.AddConfigPath(THEME_USED)
	err = v.MergeInConfig()
	if err != nil {
		log.Fatal(err)
	}

	config.Theme = NewTheme()

	PARSE_STATE = true
}

type Config struct {
	Article   []*Model
	Customize []*Model
	Document  []*Model
	Page      []*Model
	I18N      []*Model
	Static    []*Model
	Site      *Site
	Theme     *Theme
}

func NewConfig() *Config {
	return config
}
