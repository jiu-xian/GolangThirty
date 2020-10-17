package parser

import (
	"learngo/GolangThirty/crawler/engine"
	"regexp"
)

const CityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(CityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, "User: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}
	return result
}
