package parse

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/pkg/util"
)

type Path struct{}

func NewPath() *Path {
	return &Path{}
}

func (p *Path) Route(lang, path, slug string) string {
	dir := filepath.Dir(path)
	name := util.NewPath().Name(path)
	if slug != "" {
		name = slug
	}

	path = filepath.Join(lang, dir, name)
	return fmt.Sprintf("/%s/", p.format(path))
}

func (p *Path) Archive(lang, archive string) string {
	path := filepath.Join(lang, conf.META_ARCHIVE, archive)
	return fmt.Sprintf("/%s/", p.format(path))
}

func (p *Path) Category(lang, category string) string {
	path := filepath.Join(lang, conf.META_CATEGORY, category)
	return fmt.Sprintf("/%s/", p.format(path))
}

func (p *Path) Tag(lang, tag string) string {
	path := filepath.Join(lang, conf.META_TAG, tag)
	return fmt.Sprintf("/%s/", p.format(path))
}

func (p *Path) Catalog(lang, catalog string) string {
	path := filepath.Join(lang, catalog)
	return fmt.Sprintf("/%s/", p.format(path))
}

func (p *Path) format(path string) string {
	path = strings.ReplaceAll(path, " ", "-")
	return strings.ToLower(path)
}
