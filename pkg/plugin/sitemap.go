package plugin

import "github.com/beevik/etree"

type Sitemap struct {
	Loc        string // localtion
	Lastmod    string // last update
	Changefreq string // change freq
	Priority   string // priority
}

func NewSitemap() *Sitemap {
	return &Sitemap{}
}

func (s *Sitemap) List(args []*Sitemap) (string, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	sitemapindex := doc.CreateElement("sitemapindex")
	sitemapindex.CreateAttr("xmlns", "http://www.sitemaps.org/schemas/sitemap/0.9")

	for _, v := range args {
		sitemap := sitemapindex.CreateElement("sitemap")
		if v.Loc != "" {
			loc := sitemap.CreateElement("loc")
			loc.CreateText(v.Loc)
		}

		if v.Lastmod != "" {
			lastmod := sitemap.CreateElement("lastmod")
			lastmod.CreateText(v.Lastmod)
		}
	}

	doc.Indent(4)
	return doc.WriteToString()
}

func (s *Sitemap) Item(args []*Sitemap) (string, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	urlset := doc.CreateElement("urlset")
	urlset.CreateAttr("xmlns", "http://www.sitemaps.org/schemas/sitemap/0.9")

	for _, v := range args {
		url := urlset.CreateElement("url")
		if v.Loc != "" {
			loc := url.CreateElement("loc")
			loc.CreateText(v.Loc)
		}

		if v.Lastmod != "" {
			lastmod := url.CreateElement("lastmod")
			lastmod.CreateText(v.Lastmod)
		}

		if v.Changefreq != "" {
			changefreq := url.CreateElement("changefreq")
			changefreq.CreateText(v.Changefreq)
		}

		if v.Priority != "" {
			priority := url.CreateElement("priority")
			priority.CreateText(v.Priority)
		}
	}

	doc.Indent(4)
	return doc.WriteToString()
}
