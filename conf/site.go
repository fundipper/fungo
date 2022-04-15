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
		Image        struct {
			Source    string
			Target    string
			Attribute map[string]string
		}
		Video []struct {
			Host string
			Path string
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
