package html

type Node interface {
	TokenLiteral() string
	String() string
}

type InlineElement interface {
	Node
	inlineNode()
}

type BlockElement interface {
	Node
	blockNode()
}
