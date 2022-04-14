package conf

import (
	"log"
)

// Site site
type Site struct {
	Origin   string
	Port     string
	Name     string
	Slogan   string
	Theme    string
	Markdown struct {
		Highlighting string
		Lazyload     struct {
			Class struct {
				Key   string
				Value string
			}
			Data struct {
				Key   string
				Value string
			}
		}
	}
	Sitemap struct {
		Changefreq string
		Priority   string
	}
	Feeds struct {
		Action  string
		Content bool
	}
}

func NewSite() *Site {
	data := Site{}
	err := v.UnmarshalKey(MODEL_SITE, &data)
	if err != nil {
		log.Fatal(err)
	}
	return &data
}
