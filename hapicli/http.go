package hapicli

import (
	"errors"
	"strings"
)

var (
	errorMethod = errors.New("Method must be one of GET, POST, PUT, PATCH or DELETE.")
)

// AbstractRequester is an interface that wrapp different types of method for the Request
type AbstractRequester interface {
	Method() string
	UrlVariables() []string
	MessageBody() string
	Headers() string
}

// Request is an object that contain the http.Request
type Request struct {
	method       string
	urlVariables []string
	messageBody  string
	headers      string
}

// New create a Request
// - param method		string		GET, POST, PUT, PATCH or DELETE
// - param urlVariables	[]string	The value of the URL variables contained in the URL template
// - param messageBody	string	The messageBody to send with the request
// - param headers		string		Optional headers
func NewRequest(method string, urlVariables []string, messageBody string, headers string) (r *Request, err error) {
	method = strings.ToUpper(method)
	// init variables for checking the method
	mValid := false
	methodList := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	for _, m := range methodList {
		if m == method {
			mValid = true
		}
	}
	// if the method is not valid send an error
	if !mValid {
		return nil, errorMethod
	}

	r = &Request{
		method:       method,
		urlVariables: urlVariables,
		messageBody:  messageBody,
		headers:      headers,
	}

	//print(r)
	// return the Request and no error
	return r, nil
}

// Method will return a string
// GET, POST, PUT, PATCH or DELETE
func (r *Request) Method() string {
	return r.method
}

// UrlVariables returns list of url variables
// The value of the URL variables is contained in the URL template.
func (r *Request) UrlVariables() []string {
	return r.urlVariables
}

// MessageBody returns The message body to be sent with the request.
func (r *Request) MessageBody() string {
	return r.messageBody
}

// Headers return The optional headers.
func (r *Request) Headers() string {
	return r.headers
}
