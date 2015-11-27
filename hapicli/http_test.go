package hapicli

import (
	"testing"
)

var dataRequest = []struct {
	Title          string
	InMethod       string
	InUrlVariables []string
	InMessageBody  string
	InHeaders      string
	OutError       error
}{
	{"A", "GET", []string{"one", "tow"}, "message body", "in headers", nil},
	{"B", "POST", []string{"one"}, "", "", nil},
	{"C", "PUT", []string{""}, "", "", nil},
	{"D", "PATCH", []string{"", ""}, "", "", nil},
	{"E", "DELETE", []string{""}, "", "", nil},
	{"F", "", []string{""}, "", "", errorMethod},
	{"G", "TEG", []string{""}, "", "", errorMethod},
	{"H", "ET P", []string{""}, "", "", errorMethod},
}

func TestNewRequest(t *testing.T) {
	for _, v := range dataRequest {
		_, err := NewRequest(v.InMethod, v.InUrlVariables, v.InMessageBody, v.InHeaders)
		// check errors
		if err != v.OutError {
			t.Error("for test ", v.Title, "expected ", v.OutError, "got", err)
		}
	}
}

func TestMethod(t *testing.T) {
	for _, v := range dataRequest {
		r, err := NewRequest(v.InMethod, v.InUrlVariables, v.InMessageBody, v.InHeaders)
		if err != nil {
			t.Skip()
		}
		// call Method
		if m := r.Method(); m != v.InMethod {
			t.Error("for test ", v.Title, "expected ", v.InMethod, "got", m)
		}
	}
}

func TestUrlVariables(t *testing.T) {
	for _, v := range dataRequest {
		r, err := NewRequest(v.InMethod, v.InUrlVariables, v.InMessageBody, v.InHeaders)
		if err != nil {
			t.Skip()
		}
		// call UrlVariables
		uv := r.UrlVariables()
		if len(uv) != len(v.InUrlVariables) {
			t.Error("for test ", v.Title, "expected ", v.InUrlVariables, "got", uv)
		}
		for key, w := range v.InUrlVariables {
			if uv[key] != w {
				t.Error("for test ", v.Title, "expected ", w, "got", uv[key])
			}
		}
	}
}

func TestMessageBody(t *testing.T) {
	for _, v := range dataRequest {
		r, err := NewRequest(v.InMethod, v.InUrlVariables, v.InMessageBody, v.InHeaders)
		if err != nil {
			t.Skip()
		}
		// call MessageBody
		if mb := r.MessageBody(); mb != v.InMessageBody {
			t.Error("for test ", v.Title, "expected ", v.InMessageBody, "got", mb)
		}
	}
}

func TestHeaders(t *testing.T) {
	for _, v := range dataRequest {
		r, err := NewRequest(v.InMethod, v.InUrlVariables, v.InMessageBody, v.InHeaders)
		if err != nil {
			t.Skip()
		}
		// call Headers
		if h := r.Headers(); h != v.InHeaders {
			t.Error("for test ", v.Title, "expected ", v.InHeaders, "got", h)
		}
	}
}
