package plugin

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/pkg/util"
)

var t = template.New(conf.THEME_ROOT)

func init() {
	if !conf.PARSE_STATE {
		return
	}

	option := template.FuncMap{
		"toUpper": strings.ToUpper,
		"toLower": strings.ToLower,
		"toJSON": func(v interface{}) string {
			body, err := json.Marshal(v)
			if err != nil {
				return err.Error()
			}
			return string(body)
		},
		"add": func(a, b int) int {
			return a + b
		},
		"subtract": func(a, b int) int {
			return a - b
		},
		"multiply": func(a, b int) int {
			return a * b
		},
		"divide": func(a, b int) int {
			return a / b
		},
		"remainder": func(a, b int) int {
			return a % b
		},
		"safe": func(s string) interface{} {
			return template.HTML(s)
		},
		"i18n": func(lang, s string) string {
			return NewI18N().Parse(lang, s)
		},
		"path": func(s string) string {
			return util.NewPath().Dir(s)
		},
		"slug": func(s string) string {
			return util.NewPath().Name(s)
		},
	}

	var err error
	t, err = NewHTML().Template.Funcs(option).ParseGlob(conf.THEME_HTML)
	if err != nil {
		log.Fatal(err)
	}
}

type HTML struct {
	Template *template.Template
}

func NewHTML() *HTML {
	return &HTML{
		Template: t,
	}
}

func (h *HTML) Render(w io.Writer, name string, data interface{}) error {
	return h.Template.ExecuteTemplate(w, name, data)
}

func (h *HTML) Export(path, name string, data interface{}) error {
	path = fmt.Sprintf(conf.PUBLIC_HTML, path)
	err := util.NewTree().MkDir(filepath.Dir(path))
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	return h.Template.ExecuteTemplate(f, name, data)
}
