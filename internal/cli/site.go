package cli

import (
	"path/filepath"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/pkg/util"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Site struct {
	Name string
}

func NewSite(name string) *Site {
	return &Site{
		Name: name,
	}
}

func (s *Site) Create() error {
	work, err := util.NewPath().Work()
	if err != nil {
		return err
	}
	path := filepath.Join(work, s.Name)

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
