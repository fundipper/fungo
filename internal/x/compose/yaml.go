package compose

import (
	"time"

	"github.com/fundipper/fungo/conf"
	"gopkg.in/yaml.v2"
)

type YAML struct{}

func NewYAML() *YAML {
	return &YAML{}
}

func (y *YAML) Generate(source []byte, name string) ([]byte, error) {
	data := make(map[string]interface{})
	err := yaml.Unmarshal(source, data)
	if err != nil {
		return nil, err
	}

	data[conf.META_TITLE] = name
	data[conf.META_DATE] = time.Now().Format("2006-01-02")
	return yaml.Marshal(data)
}
