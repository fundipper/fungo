package message

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
	"github.com/fundipper/fungo/pkg/util"
	"github.com/julienschmidt/httprouter"
)

type Sitemap struct{}

func NewSitemap() *Sitemap {
	return &Sitemap{}
}

func (s *Sitemap) Serve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := util.NewPath().Name(r.RequestURI)

	result, err := compose.NewSitemap().Item(name)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, result)
}

func (s *Sitemap) Build(name string) (int, error) {
	path := fmt.Sprintf(conf.PUBLIC_XML, strings.ToLower(name))
	result, err := compose.NewSitemap().Item(name)
	if err != nil {
		return 0, err
	}
	return util.NewTree().WriteFile(path, []byte(result))
}

func (s *Sitemap) ServeList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result, err := compose.NewSitemap().List()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, result)
}

func (s *Sitemap) BuildList() (int, error) {
	result, err := compose.NewSitemap().List()
	if err != nil {
		return 0, err
	}
	return util.NewTree().WriteFile(conf.PUBLIC_SITEMAP, []byte(result))
}
