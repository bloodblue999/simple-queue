package router

import (
	"net/http"

	"github.com/adjust/rmq/v5"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Queue         rmq.Queue
	Server        *echo.Echo
	ResultChannel chan string
}

func (r *Router) Handler(ctx echo.Context) error {
	if err := r.Queue.PublishBytes([]byte{}); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	generatedUUID := <-r.ResultChannel

	return ctx.String(http.StatusOK, generatedUUID)
}
