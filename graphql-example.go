package main

/*
go get -v github.com/graph-gophers/graphql-go

test:
  curl -d '{"query": "query GetAllLinks { links(id: \"1000\") { id url } }" }' localhost:8000/query | jq
*/

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var s = `
type Link {
	id: ID!
	url: String!
}

type Query {
	links(id: ID!): [Link!]!
}
`

var links = []Link{
	{"1000", "https://blog.bin3xish477.com"},
	{"1001", "https://aws.amazon.com"},
	{"1002", "https://cloudflare.com"},
	{"1003", "https://protonmail.com"},
  {"1004", "https://spacex.com"},
}

type query struct{}

func (_ *query) Links(args LinkArgs) []Link {
	var lnks []Link
	for _, lnk := range links {
		if string(lnk.ID()) == args.Id {
			lnks = append(lnks, lnk)
		}
	}
	return lnks
}

type Link struct {
	IdField  string
	UrlField string
}

type LinkArgs struct {
	Id string
}

func (l Link) ID() graphql.ID {
	return graphql.ID(l.IdField)
}

func (l Link) Url() string {
	return l.UrlField
}

func main() {
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{
		Schema: schema,
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
