package cli

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/parse"
	"github.com/fundipper/fungo/pkg/cache"
	"github.com/fundipper/fungo/pkg/util"
)

type Watch struct{}

func NewWatch() *Watch {
	once.Do(func() {
		NewPrase().Run()
	})
	return &Watch{}
}

func (w *Watch) Run() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					code, err := util.NewPath().Code()
					if err != nil {
						log.Fatal(err)
					}
					err = syscall.Exec(code, os.Args, os.Environ())
					if err != nil {
						log.Fatal(err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	_ = watcher.Add(conf.ROOT_TOML)
	_ = watcher.Add(conf.THEME_TOML)
	_ = watcher.Add(conf.THEME_I18N)
	_ = watcher.Add(conf.THEME_TEMPLATES)

	for _, v := range conf.NewConfig().Static {
		for _, item := range v.Subtree {
			path := conf.THEME_ASSETS
			if !strings.HasSuffix(path, v.Name) {
				path = conf.CONTENT_MEDIA
			}

			path = fmt.Sprintf("%s/%s", path, item)
			_ = watcher.Add(path)
		}
	}

	for _, v := range conf.NewConfig().Article {
		route, ok := cache.NewSet().Get(v.Name)
		if !ok {
			continue
		}
		for item := range route {
			key := parse.NewKey().Path(item)
			path, ok := cache.NewString().Get(key)
			if ok {
				_ = watcher.Add(path)
			}
		}
	}

	for _, v := range conf.NewConfig().Document {
		route, ok := cache.NewSet().Get(v.Name)
		if !ok {
			continue
		}
		for item := range route {
			key := parse.NewKey().Path(item)
			path, ok := cache.NewString().Get(key)
			if ok {
				_ = watcher.Add(path)
			}
		}
	}

	for _, v := range conf.NewConfig().Page {
		route, ok := cache.NewSet().Get(v.Name)
		if !ok {
			continue
		}
		for item := range route {
			key := parse.NewKey().Path(item)
			path, ok := cache.NewString().Get(key)
			if ok {
				_ = watcher.Add(path)
			}
		}
	}

	<-done
}
