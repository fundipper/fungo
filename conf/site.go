package conf

type (
	// Site site
	Site struct {
		Origin   string
		Port     string
		Name     string
		Slogan   string
		Theme    string
		Size     int
		Amount   int
		Markdown struct {
			State struct {
				GFM            bool
				Table          bool
				Strikethrough  bool
				Linkify        bool
				TaskList       bool
				Emoji          bool
				DefinitionList bool
				Footnote       bool
				Typographer    bool
				Mathjax        bool
				Mermaid        bool
				Fences         bool
			}
			Highlighting struct {
				State      bool
				Theme      string
				LineNumber bool
			}
			Image struct {
				State     bool
				Source    string
				Target    string
				Attribute map[string]string
			}
			Link struct {
				State     bool
				Source    []string
				Attribute map[string]string
			}
			Video struct {
				State     bool
				Source    []*Embed
				Attribute map[string]string
			}
		}
		Sitemap struct {
			Changefreq string
			Priority   string
		}
		Feeds struct {
			Action  string
			Limit   int
			Content bool
		}
	}

	Embed struct {
		Name string
		Path string
	}
)
