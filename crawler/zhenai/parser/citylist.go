package parser

import (
	"golang_learn/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9A-Z]*)" [^>]*>([^<]+)</a>`
func ParseCityList(content []byte) engine.ParseResult{
	//                          <a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f>阿坝</a>
	//							<a href="http://www.zhenai.com/zhenghun/[a-z0-9A-Z]*" [^>]*>[^<]+</a>
	e,_ :=regexp.Compile(cityListRe)

	matches := e.FindAllSubmatch(content,-1)

	result := engine.ParseResult{}
	limit := 2
	for _,m := range matches {

		result.Items = append(result.Items,"City " +string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url : string(m[1]),
				ParseFunc : ParseCity,
			})
		limit--
		if(limit == 0) {
			break;
		}
	}
	return result
}
