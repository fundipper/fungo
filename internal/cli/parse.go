package cli

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/fatih/color"
	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/parse"
	"github.com/fundipper/fungo/pkg/util"
)

type Parse struct{}

func NewPrase() *Parse {
	if !conf.PARSE_STATE {
		log.Fatal(conf.ERROR_CONFIG)
	}
	color.Cyan(_BANNER, _VERSION)
	return &Parse{}
}

func (p *Parse) Run() {
	var wg sync.WaitGroup

	for _, v := range conf.NewConfig().Article {
		path := filepath.Join(conf.CONTENT_ROOT, v.Name)
		result, err := util.NewTree().ReadDir(path)
		if err != nil {
			continue
		}
		for _, item := range result {
			wg.Add(1)
			go func(model *conf.Model, path string) {
				defer wg.Done()

				_ = parse.NewMarkdown(model).Parse(path)
			}(v, item)
		}
	}

	for _, v := range conf.NewConfig().Document {
		path := filepath.Join(conf.CONTENT_ROOT, v.Name)
		result, err := util.NewTree().ReadDir(path)
		if err != nil {
			continue
		}
		for _, item := range result {
			wg.Add(1)
			go func(model *conf.Model, path string) {
				defer wg.Done()

				_ = parse.NewMarkdown(model).Parse(path)
			}(v, item)
		}
	}

	for _, v := range conf.NewConfig().Page {
		path := filepath.Join(conf.CONTENT_ROOT, v.Name)
		result, err := util.NewTree().ReadDir(path)
		if err != nil {
			continue
		}
		for _, item := range result {
			wg.Add(1)
			go func(model *conf.Model, path string) {
				defer wg.Done()

				_ = parse.NewMarkdown(model).Parse(path)
			}(v, item)
		}
	}

	wg.Wait()
}
