package views

import (
	"bytes"
	"testing"
)

func TestParseTemplateDir(t *testing.T) {
	tmpl, err := ParseTemplateDir("./templates")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", tmpl.DefinedTemplates())
	t.Logf("%s\n", tmpl.Name())
}

func TestRenderer_Render(t *testing.T) {
	r := MustNewRenderer("./templates")

	type args struct {
		name string
		data any
	}
	tests := []struct {
		name    string
		r       Renderer
		args    args
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
			wr := &bytes.Buffer{}
			if err := tt.r.Render(wr, tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Renderer.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("%s\n", wr.String())
		})
	}
}
