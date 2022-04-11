package cli

import (
	"errors"
	"fmt"

	"github.com/fundipper/fungo/conf"
	"github.com/fundipper/fungo/internal/x/compose"
	"github.com/fundipper/fungo/pkg/util"
)

type File struct {
	Model string
}

func NewFile(model string) *File {
	return &File{
		Model: model,
	}
}

func (f *File) Create(name string) error {
	dest := fmt.Sprintf(conf.CONTENT_MD, name)
	if util.NewPath().Exist(dest) {
		return errors.New(conf.ERROR_EXIST)
	}

	source := fmt.Sprintf(conf.SOURCE_YAML, f.Model)
	body, err := util.NewTree().ReadFile(source)
	if err != nil {
		return err
	}
	body, err = compose.NewYAML().Generate(body, name)
	if err != nil {
		return err
	}

	data := append([]byte{45, 45, 45, 10}, body...)
	data = append(data, 45, 45, 45)

	_, err = util.NewTree().WriteFile(dest, data)
	return err
}
