package engine

import (
	"golang_learn/crawler/fetcher"
	"log"
)

func Run(seeds ...Request){
	var requests []Request
	for _, r := range seeds {
		requests = append(requests,r)
	}

	for len(requests) >0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf(" Featching url: %s",r.Url)

		body,err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf(" Fetcher error , url: %s , err: %v",r.Url,err)
			continue
		}

		parseResult := r.ParseFunc(body)
		requests = append(requests,parseResult.Requests...)

		for _,itme := range parseResult.Items {
			log.Printf(" get item %s",itme)
		}


	}
}
