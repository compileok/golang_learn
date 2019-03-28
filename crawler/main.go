package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	resp,err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	all,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s",all)

	//process(all)
	printCityList(all)
}


func process(all string) {

}

func printCityList(content []byte){
	//                          <a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f>阿坝</a>
	//							<a href="http://www.zhenai.com/zhenghun/[a-z0-9A-Z]*" [^>]*>[^<]+</a>
	e,_ :=regexp.Compile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9A-Z]*)" [^>]*>([^<]+)</a>`)
	/*matches := e.FindAll(content,-1)
	for _,m := range matches {
		fmt.Printf("%s\n", m)
	}
	fmt.Printf("%d", len(matches))*/

	matches := e.FindAllSubmatch(content,-1)
	for _,m := range matches {
		/*for _, subMatch := range m {
			//fmt.Printf("city: %s , url: %s \n", subMatch[2],subMatch[1])
			//fmt.Printf(" %s",subMatch)

		}*/
		fmt.Printf("url: %s , name: %s",m[1],m[2])
		fmt.Println()
	}
}