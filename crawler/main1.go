package main

import (
	"golang_learn/crawler/engine"
	"golang_learn/crawler/zhenai/parser"
)

func main(){
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
