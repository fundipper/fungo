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

	CONTENT_ROOT  = "content"
	CONTENT_MEDIA = "content/media"
	CONTENT_MD    = "content/%s.md"

	MESSAGE_CONTENT  = "content"
	MESSAGE_PATH     = "path"
	MESSAGE_TOC      = "toc"
	MESSAGE_PAGE     = "page"
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
	MODEL_THMEM     = "theme"

	FEEDS_ATOM = "atom.xml"
	FEEDS_RSS  = "rss.xml"
	FEEDS_JSON = "feeds.json"

	PUBLIC_ROOT    = "public"
	PUBLIC_FEEDS   = "public/%s"
	PUBLIC_ASSETS  = "public/assets"
	PUBLIC_MEIDA   = "public/media"
	PUBLIC_SITEMAP = "public/sitemap.xml"
	PUBLIC_HTML    = "public%sindex.html"
	PUBLIC_XML     = "public/sitemap/%s.xml"

	ROOT_TOML      = "config.toml"
	PREFIX_CONTENT = "content/"
	PREFIX_PAGE    = "content/%s/"

	SITEMAP_XML = "%s/sitemap/%s.xml"
	SOURCE_YAML = "source/%s.yaml"
	SUFFIX_MD   = ".md"

	THEME_ROOT    = "theme"
	THEME_DEFAULT = "fungo"

	URL_SITE  = "https://github.com/fundipper/site"
	URL_THEME = "https://github.com/fundipper/theme"

	ERROR_CONFIG  = "config is nil"
	ERROR_CONTENT = "content is nil"
	ERROR_META    = "meta is nil"
	ERROR_TOC     = "toc is nil"
	ERROR_CATALOG = "catalog is nil"
	ERROR_EXIST   = "file is exist"
	ERROR_MATCH   = "name not match"
	ERROR_RANK    = "rank is nul"
)

var (
	v               *viper.Viper
	config          Config
	PARSE_STATE     bool
	THEME_USED      string
	THEME_I18N      string
	THEME_ASSETS    string
	THEME_TEMPLATES string
	THEME_HTML      string
	THEME_TOML      string
)

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

func init() {
	v = viper.New()
	v.SetConfigName(CONFIG_ROOT)
	v.SetConfigType(CONFIG_TYPE)
	v.AddConfigPath("./")

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("read config error: %v", err)
		return
	}

	err = v.Unmarshal(&config)
	if err != nil {
		log.Printf("pasre config error: %v", err)
		return
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
		log.Printf("merge config error: %v", err)
		return
	}

	theme := Theme{}
	err = v.UnmarshalKey(MODEL_THMEM, &theme)
	if err != nil {
		return
	}
	config.Theme = &theme

	PARSE_STATE = true
}

func NewConfig() *Config {
	return &config
}
