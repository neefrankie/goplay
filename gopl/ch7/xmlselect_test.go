package ch7

import (
	"strings"
	"testing"
)

func TestXMLSelect(t *testing.T) {
	file := `
<html>
	<body>
		<div><div><h2>1 Introduction</h2></div></div>
		<div><div><h2>2 Documents</h2></div></div>
		<div><div><h2>3 Logical Structure</h2></div></div>
		<div><div><h2>4 Physical Structures</h2></div></div>
		<div><div><h2>5 Conformance</h2></div></div>
		<div><div><h2>6 Notation</h2></div></div>
		<div><div><h2>A reference</h2></div></div>
		<div><div><h2>B Definitions for Character Normalization</h2></div></div>
	</body>
</html>
`
	err := XmlSelect(strings.NewReader(file), []string{"div", "div", "h2"}, func(data string) {
		t.Logf("%s", data)
	})

	if err != nil {
		t.Error(err)
	}
}
