package conf

type (
	Model struct {
		Name     string
		Route    string
		Path     string
		Action   string
		Template string
		Catalog  bool
		Contents bool
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
