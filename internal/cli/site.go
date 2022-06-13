package cli

import (
	"fmt"
	"path/filepath"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/pkg/util"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Site struct{}

func NewSite() *Site {
	return &Site{}
}

func (s *Site) Create(name string) error {
	work, err := util.NewPath().Work()
	if err != nil {
		return err
	}

	path := filepath.Join(work, name)
	fmt.Printf("clone site to %s\n", path)

	err = s.Clone(path)
	if err != nil {
		return err
	}

	path = filepath.Join(work, name, conf.THEME_ROOT, conf.THEME_DEFAULT)
	fmt.Printf("clone theme to %s\n", path)

	return NewTheme().Clone(path)
}

func (s *Site) Clone(path string) error {
	r, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:               conf.URL_SITE,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		ReferenceName:     plumbing.ReferenceName("refs/heads/master"),
	})
	if err != nil {
		return err
	}

	ref, err := r.Head()
	if err != nil {
		return err
	}

	_, err = r.CommitObject(ref.Hash())
	return err
}
