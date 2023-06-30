package parser

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestHTMLToDoc(t *testing.T) {
	type args struct {
		inputHTMLLikeString string
	}

	const successHTML = "\\nsuccess\\n"
	const failureHTML = "<script>test</script>"

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				inputHTMLLikeString: successHTML,
			},
			want:    " success ",
			wantErr: false,
		},
		{
			name: "failure",
			args: args{
				inputHTMLLikeString: failureHTML,
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HTMLToDoc(tt.args.inputHTMLLikeString)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTMLToDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTMLToDoc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractText(t *testing.T) {
	const testSuccessHTMLNode = "<p>test</p>"
	const testFailureHTMLNode = "<script>test<//////script>"
	const testScriptHTMLNode = "<script>test</script>"
	const testLineBreakHTMLNode = `<p>test\ntest</p>`

	successDoc, err := html.Parse(strings.NewReader(testSuccessHTMLNode))
	if err != nil {
		t.Errorf("failed to parse html: %v", err)
	}

	failureDoc, err := html.Parse(strings.NewReader(testFailureHTMLNode))
	if err != nil {
		t.Errorf("failed to parse html: %v", err)
	}

	scriptDoc, err := html.Parse(strings.NewReader(testScriptHTMLNode))
	if err != nil {
		t.Errorf("failed to parse html: %v", err)
	}

	lineBreakDoc, err := html.Parse(strings.NewReader(testLineBreakHTMLNode))
	if err != nil {
		t.Errorf("failed to parse html: %v", err)
	}

	type args struct {
		n *html.Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				n: successDoc,
			},
			want: "test",
		},
		{
			name: "line break",
			args: args{
				n: lineBreakDoc,
			},
			want: "test test",
		},
		{
			name: "failure",
			args: args{
				n: failureDoc,
			},
			want: "",
		},
		{
			name: "script",
			args: args{
				n: scriptDoc,
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractText(tt.args.n); got != tt.want {
				t.Errorf("extractText() = %v, want %v", got, tt.want)
			}
		})
	}
}
