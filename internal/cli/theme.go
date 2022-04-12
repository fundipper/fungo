package cli

import (
	"path/filepath"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/pkg/util"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Theme struct{}

func NewTheme() *Theme {
	return &Theme{}
}

func (t *Theme) Create(name string) error {
	work, err := util.NewPath().Work()
	if err != nil {
		return err
	}

	path := filepath.Join(work, conf.THEME_ROOT, name)
	return t.Clone(path)
}

func (t *Theme) Clone(path string) error {
	r, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:               conf.URL_THEME,
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
