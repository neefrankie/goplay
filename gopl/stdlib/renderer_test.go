package stdlib

import (
	"testing"
)

type TemplateData struct {
	SITENAME string
	SITEURL  string
	CSRF     string
}

var data = TemplateData{
	SITENAME: "Theory and Practice",
	SITEURL:  "https://siongui.github.io/",
	CSRF:     MustRandHex(8),
}

func TestParseFSTemplateDir(t *testing.T) {
	tmpl, err := ParseFSTemplateDir(templates, "templates")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", tmpl.DefinedTemplates())
	t.Logf("%s\n", tmpl.Name())
}

func TestRenderer_RenderString(t *testing.T) {

	r := MustNewFSRenderer(templates, "templates", "templates/layout")

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
		{
			name: "index.html",
			r:    r,
			args: args{
				name: "index.html",
				data: nil,
			},
			wantErr: false,
		},
		{
			name: "upload.html",
			r:    r,
			args: args{
				name: "upload.html",
				data: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.RenderString(tt.args.name, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Renderer.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			_, err = SaveString("build/"+tt.name, got)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
