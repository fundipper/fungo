package parse

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fundipper/fungo/conf"
)

type Path struct{}

func NewPath() *Path {
	return &Path{}
}

func (p *Path) Dir(path string) string {
	dir := strings.Split(path, "/")
	return filepath.Join(dir[1 : len(dir)-1]...)
}

func (p *Path) Route(lang, dir, slug string) string {
	path := filepath.Join(lang, dir, slug)
	return fmt.Sprintf("/%s/", p.format(path))
}

func (p *Path) Contents(lang, path string) string {
	path = filepath.Join(lang, p.Dir(path))

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

func (p *Path) Page(path string, page int) string {
	if page != 0 {
		return fmt.Sprintf("%spage/%d/", path, page)
	}
	return path
}

func (p *Path) format(path string) string {
	path = strings.ReplaceAll(path, " ", "-")
	return strings.ToLower(path)
}
