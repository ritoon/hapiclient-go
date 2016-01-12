package hal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// The Resource Object described in the
// JSON Hypertext Application Language (draft-kelly-json-hal-07)
// see https://tools.ietf.org/html/draft-kelly-json-hal-07#section-4
//
// Note:	When trying to find a Link or an embedded Resource
//			by their relation type (Rel), the search is done by
//			comparing the lower-case relation name.
//
// "When extension relation types are compared, they MUST be compared as
// strings [...] in a case-insensitive fashion."
// see https://tools.ietf.org/html/rfc5988#section-4.2
type resource struct {
	state             []string
	links             map[string]link
	embeddedResources map[string]string
}

// NewResource is creating a resource or an error if some params are nil
func NewResource(st []string, ls map[string]link, er map[string]string) (*resource, error) {
	if len(st) == 0 || len(ls) == 0 || len(er) == 0 {
		return nil, errors.New("Hal: please fill all params")
	}
	r := resource{st, ls, er}
	return &r, nil
}

// All the properties of the resource
// ("_links" and "_embedded" not included).
// return	[]string
func (r *resource) State() []string {
	return r.state
}

// All the links directly available in the resource.
// The key is the relation type (Rel) and the value
// can be either a Link or a numeric array of Links.
//
// Note that there is no guarantees as to the order of the links.
// return	[]link
func (r *resource) AllLinks() map[string]link {
	return r.links
}

// Finds a unique embedded resource by its relation type.
// param $rel	RegisteredRel|CustomRel		The relation type.
// return	Resource	The Resource referenced by the given rel.
// throws EmbeddedResourceNotUniqueException
// throws RelNotFoundException
func (r *resource) AllEmbeddedResources() (map[string]string, error) {
	if len(r.embeddedResources) == 0 {
		return nil, errors.New("Hal: there is no embedded Resources")
	}
	return r.embeddedResources, nil
}

// Finds a unique link by its relation type.
// param $rel	RegisteredRel|CustomRel		The relation type.
// return	Link	The Link referenced by the given rel.
// throws LinkNotUniqueException
// throws RelNotFoundException
func (r *resource) Link(rel string) (*link, error) {
	l, ok := r.links[rel]
	if !ok {
		return nil, errors.New("Hal: Rel not found")
	}
	return &l, nil
}

// Finds an array of links by their relation type.
// Note that there is no guarantees as to the order of the links.
// param $rel	RegisteredRel|CustomRel		The relation type.
// return	Numeric array of links referenced by the given rel
// throws LinkUniqueException
// throws RelNotFoundException
func (r *resource) Links(rel string) (map[string]link, error) {
	l, err := r.Link(rel)
	ls := make(map[string]link)
	ls[rel] = *l
	if err != nil {
		return nil, err
	}
	return ls, nil
}

// Finds an array of embedded resources by their relation type.
// Note that there is no guarantees as to the order of the resources.
// param $rel	RegisteredRel|CustomRel		The relation type.
// return	Numeric array of embedded resources referenced by the given rel.
// throws EmbeddedResourceUniqueException
// throws RelNotFoundException
func (r *resource) EmbeddedResources(rel string) (int, error) {
	//todo
	return 0, nil
}

// Looks for the given relation name in a case-insensitive
// fashion and returns the corresponding value.
// return	mixed	The value in $a matching the relation name
//					or null if not found.
func (r *resource) findByRel(table []string, rel string) (int, error) {
	for key, v := range table {
		s := strings.ToLower(v)
		if s == rel {
			return key, nil
		}
	}
	return 0, errors.New("Hal: the relation has not been found.")
}

// Builds a Resource from its JSON representation.
// param $json		string|array|object		A JSON representing the resource.
// return	Resource
func NewRessourcefromJson(data []byte) (*resource, error) {
	var aux struct {
		state             []string          `json:"state"`
		links             map[string]link   `json:"links"`
		embeddedResources map[string]string `json:"embeddedResources"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&aux); err != nil {
		return nil, fmt.Errorf("JSON must be a string, an array or an object : %v", err)
	}
	r := resource{aux.state, aux.links, aux.embeddedResources}
	return &r, nil
}

func (r resource) extractState(data []byte) []string {
	//todo
	return r.state
}

func (r resource) extractByRel(data []byte, rel string) ([]string, error) {
	var out []string
	if len(rel) == 0 {
		return nil, errors.New("Hal: please fill a ref relation in param")
	}

	return out, nil
}

// copyMap is copiing a map A to a map B
func copyMap(mapA map[string]string, mapB map[string]string) {
	for k, v := range mapA {
		mapB[k] = v
	}
}
