package sb

type Builder struct {
	sel []*Column
}

func (b *Builder) Select(cols []*Column) *Builder {
	b.sel = cols
	return b
}
