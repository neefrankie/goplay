package sb

type Column struct {
	name string
	as   string
}

func NewCol(n string) *Column {
	return &Column{
		name: n,
	}
}

func (c *Column) AS(n string) *Column {
	c.as = n
	return c
}

func (c *Column) String() string {
	if c.as == "" {
		return c.name
	}

	return c.name + " AS " + c.as
}
