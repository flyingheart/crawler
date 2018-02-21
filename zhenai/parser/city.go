package parser

import (
	"demo/crawler/engine"
	"regexp"
)

// <a href="http://album.zhenai.com/u/107980338" target="_blank">麦子</a>
const CityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(CityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: func(c []byte) engine.ParseResult {
			return ParseProfile(c, name)
		}})
		result.ItemType = engine.ProfileRequest
	}

	return result
}
