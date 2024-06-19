package http

import (
	"encoding/json"
	"time"

	go_http "github.com/pefish/go-http"
	"github.com/pefish/go-jsvm/module"
	"github.com/pkg/errors"
)

const ModuleName = "http_go"

type Http struct {
	vm module.IWrappedVm
}

func NewHttpModule(vm module.IWrappedVm) *Http {
	return &Http{
		vm: vm,
	}
}

type PostJsonParams struct {
	Timeout time.Duration          `json:"timeout"`
	Url     string                 `json:"url"`
	Params  interface{}            `json:"params"`
	Headers map[string]interface{} `json:"headers"`
}

func (c *Http) PostJson(params PostJsonParams) map[string]interface{} {
	if params.Timeout == 0 {
		params.Timeout = 10 * time.Second
	}
	_, bodyBytes, err := go_http.NewHttpRequester(
		go_http.WithLogger(c.vm.Logger()),
		go_http.WithTimeout(params.Timeout),
	).PostForBytes(&go_http.RequestParams{
		Url:     params.Url,
		Params:  params.Params,
		Headers: params.Headers,
	})
	if err != nil {
		c.vm.Panic(errors.Wrap(err, ""))
	}
	m := make(map[string]interface{}, 0)
	err = json.Unmarshal(bodyBytes, &m)
	if err != nil {
		c.vm.Panic(errors.Wrap(err, ""))
	}
	return m
}

type GetJsonParams struct {
	Timeout time.Duration          `json:"timeout"`
	Url     string                 `json:"url"`
	Params  interface{}            `json:"params"`
	Headers map[string]interface{} `json:"headers"`
}

func (c *Http) GetJson(params GetJsonParams) map[string]interface{} {
	if params.Timeout == 0 {
		params.Timeout = 10 * time.Second
	}
	_, bodyBytes, err := go_http.NewHttpRequester(
		go_http.WithLogger(c.vm.Logger()),
		go_http.WithTimeout(params.Timeout),
	).GetForBytes(&go_http.RequestParams{
		Url:     params.Url,
		Params:  params.Params,
		Headers: params.Headers,
	})
	if err != nil {
		c.vm.Panic(errors.Wrap(err, ""))
	}
	m := make(map[string]interface{}, 0)
	err = json.Unmarshal(bodyBytes, &m)
	if err != nil {
		c.vm.Panic(errors.Wrap(err, ""))
	}
	return m
}
