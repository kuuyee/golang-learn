package main

import (
	"golang-learn/base/crawler/engine"
	"golang-learn/base/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})

}

