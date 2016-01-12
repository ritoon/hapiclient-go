package hal

import (
	"errors"
	"strings"
)

var (
	ErrNameEmpty = errors.New("Hal: The name can't be empty")
)

// The Extension Relation Type described in:
// - section 8.2 of the HAL specification
// - section 4 of the RFC5988 - Web Linking document
// see https://tools.ietf.org/html/draft-kelly-json-hal-07#section-8.2
// see https://tools.ietf.org/html/rfc5988#section-4
type customRel struct {
	name string
}

// An Extension Relation Type SHOULD be a URI or a name
// using the CURIE syntax (prefix:reference).
// Any name as a string is accepted too.
// param s string The relation name.
func NewCustomRel(s string) (*customRel, error) {
	s, err := trimSpace(s)
	if err != nil {
		return nil, err
	}
	cr := customRel{s}
	return &cr, nil
}

// return string The relation name.
func (cr *customRel) Name() string {
	return cr.name
}

// change the name of the customRel
func (cr *customRel) SetName(s string) error {
	s, err := trimSpace(s)
	if err != nil {
		return err
	}
	cr.name = s
	return nil
}

// return string The relation name.
func (cr *customRel) String() string {
	return cr.name
}

// trimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
// this will return an error if the string is empty
func trimSpace(s string) (string, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s, ErrNameEmpty
	}
	return s, nil
}
