package parser

import (
	"golang_learn/crawler/engine"
	"golang_learn/crawler/model"
	"regexp"
	"strconv"
)

//var gender = regexp.MustCompile(``) 没有
var ageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)岁</div>`)
var marriageRe =regexp.MustCompile(`<div class="m-btn purple" [^>]*>([未婚|离异]+)</div>`)
var xinzuo = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\p{Han}]+座[^<]+)</div>`)
var height = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)[cm|CM|Cm|cM]+</div>`)
var weight = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)[kg|KG|Kg|kG]+</div>`)
var income = regexp.MustCompile(`<div class="m-btn purple" [^>]*>月收入:([\S]+)</div>`)
var hukou = regexp.MustCompile(`<div class="m-btn pink" [^>]*>籍贯:([\p{Han}]+)</div>`)


func ParseProfile(content []byte,name string) engine.ParseResult{

	profile := model.Profile{}
	profile.Name = name

	// age
	age,err := strconv.Atoi(extractString(content,ageRe))
	if(err == nil) {
		profile.Age = age
	}

	// marriage
	profile.Marriage = extractString(content,marriageRe)

	// 星座
	profile.Xinzuo = extractString(content,xinzuo)

	// 身高
	h,err := strconv.Atoi(extractString(content,height))
	if err == nil {
		profile.Height = h
	}

	// 体重
	w,err := strconv.Atoi(extractString(content,weight))
	if err == nil {
		profile.Weight = w
	}

	// 月收入
	profile.Income = extractString(content,income)

	// 户口
	profile.Hokou = extractString(content,hukou)

	result := engine.ParseResult{
		Items:[]interface{} {profile},
	}


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