package httpHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"../graphqlSchema"

	"github.com/graphql-go/graphql"
)

func executeQuery(
	query string,
	schema graphql.Schema,
) *graphql.Result {
	result := graphql.Do(
		graphql.Params{
			Schema:        graphqlSchema.Schema,
			RequestString: query,
		},
	)
	if len(result.Errors) > 0 {
		fmt.Printf(
			"\nexecuteQuery() @ RequestServer.go\n"+
				"Wrong result, unexpected errors: %v\n\n",
			result.Errors,
		)
	}
	return result
}

func getGraphqlQuery(urlParam string, body string) (string, error) {
	if urlParam != "" {
		return urlParam, nil
	} else if body != "" {
		return body, nil
	}

	return "", errors.New(
		"No query found neither in URL params or body",
	)
}

type GraphqlHandler struct{}

func (handler GraphqlHandler) ServeHTTP(
	respWriter http.ResponseWriter,
	request *http.Request,
) {
	respWriter.Header().Add("Content-Type", "application/json")
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	graphqlQuery, err := getGraphqlQuery(
		request.URL.Query().Get("query"),
		string(requestBody[:]),
	)
	if err != nil {
		fmt.Println(err)
	}

	result := executeQuery(
		graphqlQuery,
		graphqlSchema.Schema,
	)
	json.NewEncoder(respWriter).Encode(result)
}
