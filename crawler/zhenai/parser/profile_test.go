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
	result := ParseProfile(contents,"test")
	profile := result.Items[0].(model.Profile)
	fmt.Println("xingzuo: "+profile.Xinzuo)
	fmt.Printf("age: %d\n" , profile.Age)
	fmt.Println("height: ",profile.Height)
	fmt.Println("weight: ",profile.Weight)
	fmt.Println("income: ",profile.Income)
	fmt.Println("hukou:" + profile.Hokou)

}
