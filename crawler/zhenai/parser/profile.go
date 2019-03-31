package parser

import (
	"golang_learn/crawler/engine"
	"golang_learn/crawler/model"
	"regexp"
	"strconv"
)


var ageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)岁</div>`)
var marriageRe =regexp.MustCompile(`<div class="m-btn purple" [^>]*>([未婚|离异]+)</div>`)



func ParseProfile(content []byte) engine.ParseResult{

	profile := model.Profile{}

	age,err := strconv.Atoi(extractString(content,ageRe))
	if(err == nil) {
		profile.Age = age
	}

	profile.Marriage = extractString(content,marriageRe)

	result := engine.ParseResult{
		Items:[]interface{} {profile},
	}

	// 其他属性

	return result

}

func extractString(contents []byte,re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >=2  {
		return string(match[1])
	} else {
		return ""
	}



}