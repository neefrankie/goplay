package stdlib

import (
	"fmt"
	"mime"
	"mime/multipart"
	"os"
	"testing"
)

func TestCreateFormFile(t *testing.T) {
	writer := multipart.NewWriter(os.Stdout)
	defer writer.Close()

	names := []string{
		"upload1.txt",
		"upload2.txt",
	}

	for i, name := range names {
		// part is an io.Writer. It simple wraps the original multipart.Writer.
		// The CreateFormFile appends header data to the writer
		// and then returns a part wrapping this writer.
		part, err := writer.CreateFormFile(fmt.Sprintf("file%d", i), name)
		if err != nil {
			t.Fatal(err)
		}

		b, err := os.ReadFile("build/" + name)
		if err != nil {
			t.Fatal(err)
		}

		_, err = part.Write(b)

		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestParseMediaType(t *testing.T) {

	type args struct {
		v string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Content-Type: text/html",
			args: args{
				v: "text/html; charset=utf-8",
			},
			wantErr: false,
		},
		{
			name: "Content-Disposition",
			args: args{
				v: `form-data; nam="myFile"; filename="test.txt"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, param, err := mime.ParseMediaType(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("mime.ParseMediaType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("mediatype: %s, param %v", typ, param)
		})
	}
}

func TestFormatMediaType(t *testing.T) {
	type args struct {
		t     string
		param map[string]string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "text/html",
			args: args{
				t: "text/html",
				param: map[string]string{
					"charset": "utf-8",
				},
			},
			want: "text/html; charset=utf-8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mime.FormatMediaType(tt.args.t, tt.args.param)

			t.Logf("%s\n", got)
		})
	}
}

func TestTypeByExtension(t *testing.T) {
	type args struct {
		ext string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: ".html",
			args: args{
				ext: ".html",
			},
		},
		{
			name: ".png",
			args: args{
				ext: ".png",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mime.TypeByExtension(tt.args.ext)

			t.Logf("%s\n", got)
		})
	}
}

func TestExtensionByType(t *testing.T) {
	type args struct {
		typ string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "html",
			args: args{
				typ: "text/html",
			},
			wantErr: false,
		},
		{
			name: "png",
			args: args{
				typ: "image/png",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mime.ExtensionsByType(tt.args.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("mime.ExtensionByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%s\n", got)
		})
	}
}
