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

func TestTextTempl(t *testing.T) {
	sweaters := Inventory{
		Material: "wool",
		Count:    17,
	}

	// The input text for a template is UTF-8-encoded text in
	// any format.
	// Actions are delimited by `{{ }}`.
	// All text outside actions is copied to the output unchanged.
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")

	if err != nil {
		t.Error(err)
	}

	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		t.Error(err)
	}
}

func parseTemplate(t *testing.T, file string) *template.Template {
	b, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	tmpl, err := template.New("test").Parse(string(b))
	if err != nil {
		t.Error(err)
	}

	return tmpl
}

func renderFile(t *testing.T, file string) {
	tmpl := parseTemplate(t, file)

	// A template may be executed directly or though
	// ExecuteTemplate, which executes an associated
	// template identified by name.
	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		t.Error(err)
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
	renderFile(t, "./templates/pipeline_var.tmpl")
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
	renderFile(t, "./templates/nested.tmpl")
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
	tmpl := parseTemplate(t, "./templates/letter.tmpl")

	for _, r := range recipients {
		err := tmpl.Execute(os.Stdout, r)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestTmplBlock(t *testing.T) {
	master := `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
	overlay := `{{define "list"}}{{join . ", "}}{{end}}`
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

	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		t.Error(err)
	}

	err = masterTmpl.Execute(os.Stdout, guardians)
	if err != nil {
		t.Error(err)
	}

	err = overlayTmpl.Execute(os.Stdout, guardians)
	if err != nil {
		t.Error(err)
	}
}

type TemplateData struct {
	SITENAME string
	SITEURL  string
}

func TestTemplateToHtml(t *testing.T) {
	data := TemplateData{
		SITENAME: "Theory and Practice",
		SITEURL:  "https://siongui.github.io/",
	}

	tmpl, err := ParseTemplateDir("templates")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s\n", tmpl.DefinedTemplates())
	t.Logf("%s\n", tmpl.Name())

	err = tmpl.ExecuteTemplate(os.Stdout, "index.html", &data)
	if err != nil {
		t.Error(err)
	}
}
