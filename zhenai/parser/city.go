package parser

import (
	"learnGo/crawier/engine"
	"regexp"
)

const cityRe = "<th><a href=\"(http://album.zhenai.com/u/[0-9a-zA-Z]+)\".*?>(.*?)</a></th>"

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)

	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Items = append(result.Items,"User " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		//fmt.Println()
	}

	return result
}
