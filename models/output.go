package models

import (
	"github.com/gsiddharth/wandr/errors"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

type Output struct {
	Result interface{}
	Status string
	Code   int
}

func NewOutput(result interface{}, status string, code int) *Output {
	return &Output{Result: result, Status: status, Code: code}
}

func NewErrorOutput(err errors.Error) *Output {
	return &Output{Status: err.Description, Code: err.Code}
}

func (self *Output) Render(rw http.ResponseWriter) {
	r := render.New(render.Options{})
	r.JSON(rw, self.Code, self)

}
