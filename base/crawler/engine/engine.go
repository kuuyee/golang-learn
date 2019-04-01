package engine

import (
	"fmt"
	"golang-learn/base/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		fmt.Printf("Seeds number is : %d\n",len(seeds))
		requests = append(requests, r)
	}

	fmt.Println(requests);

	for len(requests) > 0 {
		r := requests[0]
		//fmt.Println(r)
		requests = requests[1:]
		//fmt.Println(requests)
		log.Printf("Fetching %s",r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}
		fmt.Println(body)

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _,item :=range parseResult.Items{
			log.Printf("Got item %v",item)
		}

	}
}
