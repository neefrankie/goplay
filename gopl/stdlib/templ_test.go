package stdlib

import (
	"html/template"
	"os"
	"strings"
	"testing"
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

var data = TemplateData{
	SITENAME: "Theory and Practice",
	SITEURL:  "https://siongui.github.io/",
}

func TestParseTemplateDir(t *testing.T) {

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

func TestParseFSTemplateDir(t *testing.T) {
	tmpl, err := ParseFSTemplateDir(f, "templates")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", tmpl.DefinedTemplates())
	t.Logf("%s\n", tmpl.Name())

	err = tmpl.ExecuteTemplate(os.Stdout, "index.html", &data)
	if err != nil {
		t.Error(err)
	}
}

func TestRender(t *testing.T) {
	r := MustNewRenderer("./templates")

	h, err := r.Render("index.html", &data)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", h)
}

func TestRenderToFile(t *testing.T) {
	r := MustNewRenderer("./templates")

	f, err := os.Create("build/index.html")
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = r.RenderTo(f, "index.html", &data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderer_Render(t *testing.T) {
	type args struct {
		name string
		data any
	}
	tests := []struct {
		name    string
		r       Renderer
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Render(tt.args.name, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Renderer.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Renderer.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
