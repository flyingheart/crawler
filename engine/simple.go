package engine

import (
	"demo/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	cityTotalCount := 0
	profileTotalCount := 0
	profileDetailsTotalCount := 0

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			if parseResult.ItemType == CityRequest {
				log.Printf("[%04d]City: %v", cityTotalCount+1, item)
				cityTotalCount++
			} else if parseResult.ItemType == ProfileRequest {
				log.Printf("[%05d]Profile name : %v", profileTotalCount+1, item)
				profileTotalCount++
			} else if parseResult.ItemType == ProfileDetails {
				log.Printf("[%05d]Profile details : %v", profileDetailsTotalCount+1, item)
				profileDetailsTotalCount++
			}
		}
	}
}

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		//log.Printf("Fetcher Url "+"fetching error %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	parseResult := r.ParserFunc(body)

	return parseResult, nil
}
