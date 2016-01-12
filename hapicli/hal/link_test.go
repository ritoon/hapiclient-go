package hal

import (
	"reflect"
	"testing"
)

func TestNewLink(t *testing.T) {
	data := []struct {
		title      string
		inHref     string
		inOptional LinkOptionalParam
		outLink    link
		outErr     error
	}{
		{
			"A",
			"test",
			LinkOptionalParam{true, "pdf", "no", "the name", "profile", "title", "golang"},
			link{"test", true, "pdf", "no", "the name", "profile", "title", "golang"},
			nil,
		},

		{
			"B",
			"",
			LinkOptionalParam{},
			link{},
			ErrNameEmpty,
		},
	}
	for _, v := range data {
		out, err := NewLink(v.inHref, v.inOptional)
		if !reflect.DeepEqual(err, v.outErr) {
			t.Error("Waiting for", v.outErr, "got", err)
		}
		if err == nil && reflect.DeepEqual(out, v.outLink) {
			t.Error("Waiting for", v.outLink, "got", out)
		}
	}
}

func TestNewLinkFromJson(t *testing.T) {
	data := []struct {
		title   string
		in      []byte
		outLink link
		outErr  error
	}{
		{
			"A",
			[]byte(`{
			    "href": "http:www.greentic.com",
			    "templated": false,
			    "type": "blue",
			    "deprecation": "no",
			    "name": "fred",
			    "profile": "man",
			    "title": "title B",
			    "hreflang": "java"
  				}`),
			link{"http:www.greentic.com", false, "blue", "no", "fred", "man", "title B", "java"},
			nil,
		},
	}
	for _, v := range data {
		l, err := NewLinkFromJson(v.in)
		if !reflect.DeepEqual(err, v.outErr) {
			t.Error("Waiting for", v.outErr, "got", err)
		}

		if err == nil && reflect.DeepEqual(l, v.outLink) {
			t.Error("Waiting for", v.outLink, "got", l)
		}

	}
}
