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
		name := string(m[2])
		result.Items = append(
			result.Items,"User " + name)


		result.Requests = append(result.Requests,engine.Request{
			Url :      string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult{
				return ParseProfile(c,name)
			},
		})
	}
	return result
}
