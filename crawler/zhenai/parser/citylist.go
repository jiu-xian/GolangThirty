package parser

import (
	"learngo/GolangThirty/crawler/engine"
	"regexp"
)

const CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(CityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
