package parser

import (
	"golang_learn/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)

	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)`)
)
const cityRe = `<a href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`
func ParseCity(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityRe)
	match := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}

	for _ , m := range match {
		result.Items = append(
			result.Items,"User " + string(m[2]))


		result.Requests = append(result.Requests,engine.Request{
			Url :      string(m[1]),
			ParseFunc: engine.NilParser,
		})
	}
	return result
}
