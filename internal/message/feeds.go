package message

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
	"github.com/fundipper/fungo/pkg/util"
	"github.com/julienschmidt/httprouter"
)

type Feeds struct {
	Model string
}

func NewFeeds(model string) *Feeds {
	return &Feeds{
		Model: model,
	}
}

func (f *Feeds) Serve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := filepath.Base(r.RequestURI)

	result, err := compose.NewFeeds(f.Model).Generate(name)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, result)
}

func (f *Feeds) Build(name string) (int, error) {
	path := fmt.Sprintf(conf.PUBLIC_FEEDS, strings.ToLower(name))

	result, err := compose.NewFeeds(f.Model).Generate(name)
	if err != nil {
		panic(err)
	}
	return util.NewTree().WriteFile(path, []byte(result))
}
