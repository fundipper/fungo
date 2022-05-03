package conf

import (
	"log"
)

type (
	Model struct {
		Name     string
		Route    string
		Path     string
		Action   string
		Template string
		Catalog  bool
		Root     bool
		Sidebar  []*Node
		Subtree  []string
	}

	Node struct {
		Number string  `json:"number,omitempty"`
		Name   string  `json:"name,omitempty"`
		Route  string  `json:"route,omitempty"`
		Node   []*Node `json:"node,omitempty"`
	}
)

func NewModel() *Model {
	return &Model{}
}

func (m *Model) Article() []*Model {
	data := []*Model{}
	err := v.UnmarshalKey(MODEL_ARTICLE, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (m *Model) Customize() []*Model {
	data := []*Model{}
	err := v.UnmarshalKey(MODEL_CUSTOMIZE, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (m *Model) Document() []*Model {
	data := []*Model{}
	err := v.UnmarshalKey(MODEL_DOCUMENT, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (m *Model) Page() []*Model {
	data := []*Model{}
	err := v.UnmarshalKey(MODEL_PAGE, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (m *Model) I18N() []*Model {
	data := []*Model{}
	err := v.UnmarshalKey(MODEL_I18N, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (m *Model) Static() []*Model {
	data := []*Model{}
	err := v.UnmarshalKey(MODEL_STATIC, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
