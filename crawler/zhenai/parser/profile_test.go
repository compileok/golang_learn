package parser

import (
	"fmt"
	"golang_learn/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T){

	contents,err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s",contents)
	result := ParseProfile(contents,"test")
	profile := result.Items[0].(model.Profile)
	fmt.Println(profile.Age == 25)


}
