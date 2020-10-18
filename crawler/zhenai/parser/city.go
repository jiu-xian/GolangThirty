package parser

import (
	"learngo/GolangThirty/crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
var cityRe = regexp.MustCompile(`<span class="pager"><a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/shanghai/[^"]+)">([^<]+)</a></span>`)

func ParseCity(contents []byte) engine.ParseResult { //城市解析器

	result := engine.ParseResult{}

	matches := profileRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User: "+name) //解析结果返回ParseResult
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	matches = cityRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		pages := string(m[2])
		result.Items = append(result.Items, "Page: "+pages)
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}

	return result
}
