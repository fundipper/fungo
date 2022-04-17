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
		Size     int
		Catalog  bool
		Root     bool
		Sidebar  []*Node
		Subtree  []string
	}

	Node struct {
		Number string
		Name   string
		Route  string
		Node   []*Node
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
