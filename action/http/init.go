package http

import (
	"github.com/graphql-go/handler"
	"hcc/piccolo/action/graphql"
	"hcc/piccolo/action/http/subscription"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"net/http"
	"strconv"
)

// Init : Initialize GraphQL server
func Init() {
	logger.Logger.Println("Opening GraphQL server on port " + strconv.Itoa(int(config.GraphQL.Port)) + "...")

	var graphqlHandler = handler.New(&handler.Config{
		Schema:     &graphql.Schema,
		Pretty:     true,
		GraphiQL:   !config.GraphQL.UsePlayground,
		Playground: config.GraphQL.UsePlayground,
	})
	http.Handle("/graphql", graphqlHandler)
	logger.Logger.Println("Serving GraphQL requests on /graphql")

	http.Handle("/subscriptions", subscription.NewSubscriptionHandler())
	logger.Logger.Println("Serving GraphQL's subscription websocket requests on /subscriptions")

	err := http.ListenAndServe(":"+strconv.Itoa(int(config.GraphQL.Port)), nil)
	if err != nil {
		logger.Logger.Println(err)
		logger.Logger.Println("Failed to prepare GraphQL server!")
		return
	}
}
