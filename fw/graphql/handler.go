package graphql

import (
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type Handler http.Handler

var _ http.Handler = (*GraphGopherHandler)(nil)

type GraphGopherHandler struct {
	handler relay.Handler
}

func (g GraphGopherHandler) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	g.handler.ServeHTTP(writer, reader)
}

func NewGraphGopherHandler(api API) GraphGopherHandler {
	schema := graphql.MustParseSchema(
		api.GetSchema(),
		api.GetResolver(),
		graphql.UseStringDescriptions(),
	)
	return GraphGopherHandler{handler: relay.Handler{
		Schema: schema,
	}}
}
