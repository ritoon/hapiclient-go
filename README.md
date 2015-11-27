[![Coverage](http://gocover.io/_badge/github.com/ritoon/hapiclient-go/hapicli?0)](http://gocover.io/github.com/ritoon/hapiclient-go/hapicli)
[![Build Status](https://travis-ci.org/ritoon/hapiclient-go.svg?branch=master)](https://travis-ci.org/ritoon/hapiclient-go)
[![GoDoc](https://godoc.org/github.com/ritoon/hapiclient/hapicli?status.svg)](https://godoc.org/github.com/ritoon/hapiclient/hapicli)

Be aware that this package is still in heavy developpement. Some breaking change will occure.
Thank's for your comprehension.

# HAPI Client

An HTTP Client implementing the [HAL specification](https://tools.ietf.org/html/draft-kelly-json-hal-07).

# Requirements

Go 1.5 or higher

# Usage

Create a client:
```go
import "golang.org/x/oauth2"

pat := "mytoken"
type TokenSource struct {
    AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
    token := &oauth2.Token{
        AccessToken: t.AccessToken,
    }
    return token, nil
}

tokenSource := &TokenSource{
    AccessToken: pat,
}
oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
client := hapicli.NewClient(oauthClient)
```

Use a service:
```go

// Data for https://api.slimpay.net/alps#create-direct-debits
jsonBody := &hapicli.JsonBody{
    Amount:   "100",
    PaymentReference: "Payment 123",
    Label: "The label",
    ExecutionDate: time.Now(),
    Creditor:{
    	Reference: "democreditor"
	},
	Mandate:{
		Rum: "SLMP000001",
		Standard : &hapicli.SEPA,
		CreateSequenceType : &hapicli.FRST,
		DateSigned : time.Now(),
	},
    
}

newSepa,err := $hapicli.CreateRequest(jsonBody)

if err != nil {
    fmt.Printf("Something bad happened: %s\n\n", err)
    return err
}
```

## Versioning

Each version of the client is tagged and the version is updated accordingly.

Since Go does not have a built-in versioning, a package management tool is
recommended - a good one that works with git tags is
[gopkg.in](http://labix.org/gopkg.in).

To see the list of past versions, run `git tag`.


## Documentation

For a comprehensive list of examples, check out the [API documentation](http://www.slimpay.net/rest-hapi-crawler).

For details on all the functionality in this library, see the [GoDoc](http://godoc.org/github.com/ritoon/hapiclient-go) documentation.


## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).


