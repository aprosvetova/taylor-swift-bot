package main

import (
	"errors"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

const ApiBase = "https://api.taylor.rest"

func getTaylor() (img string, quote string, err error) {
	var res []byte

	_, res, err = fasthttp.Get(nil, ApiBase)
	if err != nil {
		return
	}
	quote = fastjson.GetString(res, "quote")

	_, res, err = fasthttp.Get(nil, ApiBase+"/image")
	if err != nil {
		return
	}
	img = fastjson.GetString(res, "url")

	if quote == "" && img == "" {
		err = errors.New("something went wrong")
	}

	return
}
