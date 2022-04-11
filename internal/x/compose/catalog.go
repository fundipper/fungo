package compose

type Catalog []*Markdown

func (c Catalog) Len() int {
	return len(c)
}

func (c Catalog) Less(i, j int) bool {
	return c[i].Date > c[j].Date
}

func (c Catalog) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
