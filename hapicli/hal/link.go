package hal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrPropMandatory = errors.New("The href property is mandatory.")
)

type link struct {
	// REQUIRED
	// Its value is either a URI [RFC3986] or a URI Template [RFC6570].<br>
	// If the value is a URI Template then the Link Object SHOULD have a
	// "templated" attribute whose value is true.
	href string

	// OPTIONAL
	// Its value is boolean and SHOULD be true when the Link Object's "href"
	// property is a URI Template.<br>
	// Its value SHOULD be considered false if it is undefined or any other
	// value than true.
	templated bool

	// OPTIONAL
	// Its value is a string used as a hint to indicate the media type
	// expected when dereferencing the target resource.
	mediaType string

	// OPTIONAL
	// Its presence indicates that the link is to be deprecated (i.e.
	// removed) at a future date.  Its value is a URL that SHOULD provide
	// further information about the deprecation.
	// A client SHOULD provide some notification (for example, by logging a
	// warning message) whenever it traverses over a link that has this
	// property.  The notification SHOULD include the deprecation property's
	// value so that a client maintainer can easily find information about
	// the deprecation.
	deprecation string

	// OPTIONAL
	// Its value MAY be used as a secondary key for selecting Link Objects
	// which share the same relation type.
	name string

	// OPTIONAL
	// Its value is a string which is a URI that hints about the profile (as
	// defined by [I-D.wilde-profile-link]) of the target resource.
	profile string

	// OPTIONAL
	// Its value is a string and is intended for labelling the link with a
	// human-readable identifier (as defined by [RFC5988]).
	title string

	// OPTIONAL
	// Its value is a string and is intended for indicating the language of
	// the target resource (as defined by [RFC5988]).
	hreflang string
}

type LinkOptionalParam struct {
	Templated   bool
	MediaType   string
	Deprecation string
	Name        string
	Profile     string
	Title       string
	Hreflang    string
}

// NewLink create a Link
// params :
// - href			string
// optional params :
// - lop LinkOptionalParam
func NewLink(href string, lop LinkOptionalParam) (*link, error) {
	href, err := trimSpace(href)
	if err != nil {
		return nil, err
	}
	l := link{href, lop.Templated, lop.MediaType, lop.Deprecation, lop.Name, lop.Profile, lop.Title, lop.Hreflang}
	return &l, nil
}

// Href get the href link property
func (l *link) Href() string {
	return l.href
}

// Templated get the templated link property
func (l *link) Templated() bool {
	return l.templated
}

// Type get the mediaType link property
func (l *link) Type() string {
	return l.mediaType
}

// Deprecation get the deprecation link property
func (l *link) Deprecation() string {
	return l.deprecation
}

// Name get the name link property
func (l *link) Name() string {
	return l.name
}

// Profile get the profile link property
func (l *link) Profile() string {
	return l.profile
}

// Title get the profile link property
func (l *link) Title() string {
	return l.title
}

// Hreflang get the hreflang link property
func (l *link) Hreflang() string {
	return l.hreflang
}

// SetHref is setting the href of link property
func (l *link) SetHref(h string) error {
	h, err := trimSpace(h)
	if err != nil {
		return err
	}
	l.href = h
	return nil
}

// SetTemplated is setting the templated of link property
func (l *link) SetTemplated(t bool) {
	l.templated = t
}

// SetType is setting the mediaType of link property
func (l *link) SetType(t string) {
	l.mediaType = t
}

// SetDeprecation is setting the deprecation of link property
func (l *link) SetDeprecation(d string) {
	l.deprecation = d
}

// SetName is setting the name of link property
func (l *link) SetName(n string) {
	l.name = n
}

// SetProfile is setting the profile of link property
func (l *link) SetProfile(p string) {
	l.profile = p
}

// SetTitle is setting the title of link property
func (l *link) SetTitle(p string) {
	l.title = p
}

// SetHreflang is setting the hreflang of link property
func (l *link) SetHreflang(h string) {
	l.hreflang = h
}

// NewLinkFromJson is creating a link from a json
// this will trow an error if the JSON is not valid
func NewLinkFromJson(data []byte) (*link, error) {
	var aux struct {
		Href        string `json:"href"`
		Templated   bool   `json:"templated"`
		MediaType   string `json:"type"`
		Deprecation string `json:"deprecation"`
		Name        string `json:"name"`
		Profile     string `json:"profile"`
		Title       string `json:"title"`
		Hreflang    string `json:"hreflang"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&aux); err != nil {
		return nil, fmt.Errorf("JSON must be a string, an array or an object : ", err)
	}
	l := link{
		aux.Href,
		aux.Templated,
		aux.MediaType,
		aux.Deprecation,
		aux.Name,
		aux.Profile,
		aux.Title,
		aux.Hreflang}
	return &l, nil
}

// String is using the interface of strings for all usages toString()
func (l *link) String() string {
	s := "(href=" + l.href

	b := "false"
	if l.templated {
		b = "true"
	}
	s += ", temlpated=" + b

	if len(l.mediaType) > 0 {
		s += ", type=" + l.mediaType
	}

	if len(l.deprecation) > 0 {
		s += ", deprecation=" + l.deprecation
	}
	if len(l.name) > 0 {
		s += ", name=" + l.name
	}
	if len(l.profile) > 0 {
		s += ", profile=" + l.profile
	}
	if len(l.title) > 0 {
		s += ", title=" + l.title
	}
	if len(l.hreflang) > 0 {
		s += ", hreflang=" + l.hreflang
	}
	s += ")"
	return s
}
