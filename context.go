package redtea

import (
	"context"
	"net/http"
)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	ResponseHeader http.Header
	Values         map[string]interface{}
	Body           []byte
	Context        context.Context
}
