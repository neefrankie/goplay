package stdlib

import (
	"os"
	"strings"
	"testing"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func TestTemplateTrivial(t *testing.T) {
	sweaters := Inventory{
		Material: "wool",
		Count:    17,
	}

	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")

	if err != nil {
		t.Error(err)
	}

	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		t.Error(err)
	}
}

// This test shows how the block and define works.
func TestTemplateBlock(t *testing.T) {
	master := `Names:
	{{block "list" .}}
		{{range .}}
		{{println "-" .}}
		{{end}}
	{{end}}`

	overlay := `{{define "list"}}{{join . ", "}}{{end}}`

	override := `{{define "list"}}Override previous list{{end}}`

	funcs := template.FuncMap{
		"join": strings.Join,
	}

	guardians := []string{
		"Gamora",
		"Groot",
		"Nebula",
		"Rocket",
		"Star-Lord",
	}

	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		t.Error(err)
	}

	t.Log("\nBlock content\n")
	err = masterTmpl.Execute(os.Stdout, guardians)
	if err != nil {
		t.Error(err)
	}

	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		t.Error(err)
	}

	t.Log("\nOverlay content\n")
	err = overlayTmpl.Execute(os.Stdout, guardians)
	if err != nil {
		t.Error(err)
	}

	// A defined name could only show up once in a parsed *Template.
	// Later definitions will overide previous definitions.
	// If you do not intend to override the list defined in overlay,
	// you should clone it:
	// overlayTmpl.Clone().Parse(override), or
	// masterTmpl.Clone().Parse(override)
	overriddenTmpl, err := overlayTmpl.Parse(override)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\nOverridden content\n")
	err = overriddenTmpl.Execute(os.Stdout, guardians)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("\nOverlay is overriden\n")
	err = overlayTmpl.Execute(os.Stdout, guardians)
	if err != nil {
		t.Fatal(err)
	}
}

const pipelineVar = `
{{"\"output\""}}

{{/* A function call */}}
{{printf "%q" "output"}}

{{/* A functional call whose final argument comes from the previous command*/}}
{{"output" | printf "%q"}}

{{/* A parenthesized argument */}}
{{printf "%q" (print "out" "put")}}

{{"put" | printf "%s%s" "out" | printf "%q"}}

{{"output" | printf "%s" | printf "%q"}}

{{/* A with action using dot */}}
{{with "output"}}{{printf "%q" .}}{{end}}

{{/* A with action that creates and uses a variable. */}}
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}

{{/* A with action that uses the variable in another action. */}}
{{with $x := "output"}}{{printf "%q" $x}}{{end}}

{{/* The same, but pipelined */}}
{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
`

func TestTmplPipelineVar(t *testing.T) {
	tmpl, err := template.New("test").Parse(pipelineVar)
	if err != nil {
		t.Error(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		t.Error(err)
	}
}

const nested = `
{{define "T1"}}
ONE
{{end}}

{{define "T2"}}
TWO
{{end}}

{{/*
Template definition must appear at the top of the template.

The define action names the template
being created by providing a string constant.
*/}}

{{define "T3"}}
{{template "T1"}} {{template "T2"}}
{{end}}

{{/*
The above defines two templates, T1 and T2,
and a third T3 that invokes the other two
when it is executed.

Finally it invokes T3
*/}}

{{template "T3"}}
`

func TestTmplNested(t *testing.T) {
	tmpl, err := template.New("test").Parse(nested)
	if err != nil {
		t.Error(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		t.Error(err)
	}
}

type Recipient struct {
	Name     string
	Gift     string
	Attended bool
}

var recipients = []Recipient{
	{
		"Aunt Mildred",
		"bone china tea set",
		true,
	},
	{
		"Uncle John",
		"moleskin pants",
		false,
	},
	{
		"Cousin Rodney",
		"",
		false,
	},
}

const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

func TestTmplLetter(t *testing.T) {
	tmpl, err := template.New("test").Parse(letter)
	if err != nil {
		t.Error(err)
	}

	for _, r := range recipients {
		err = tmpl.Execute(os.Stdout, r)
		if err != nil {
			t.Error(err)
		}
	}
}
