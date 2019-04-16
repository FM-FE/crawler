package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`

func CityList(contents []byte) engine.ParserRequests {
	re := regexp.MustCompile(cityListRe)
	// care about the [^>] and [^<], both of those define the limit of regexp
	matches := re.FindAllSubmatch(contents, -1)

	results := engine.ParserRequests{}

	for _, m := range matches {
		// results.Items = append(results.Items, "City "+string(m[2]))
		results.Requests = append(results.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: City,
		})
	}
	return results
}
