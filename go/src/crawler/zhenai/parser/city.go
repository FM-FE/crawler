package parser

import (
	"crawler/engine"
	"regexp"
)

// <tbody><tr><th><a href="http://album.zhenai.com/u/1126232797" target="_blank">泽郎登珠</a></th>
// </tr> <tr><td width="180"><span class="grayL">性别：</span>男士</td>

const Re = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)士</td>`
const NextPageRe = `<a href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)">`

func City(contents []byte) engine.ParserRequests {
	re := regexp.MustCompile(Re)
	matches := re.FindAllSubmatch(contents, -1)

	results := engine.ParserRequests{}

	for _, m := range matches { // only display the first username
		name := string(m[2])
		gender := string(m[3])
		// results.Items = append(results.Items, "User "+string(m[2]))
		results.Requests = append(results.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParserRequests {
				return Profile(c, name, gender)
			},
		})
	}

	nextPageRe := regexp.MustCompile(NextPageRe)
	nextPageMatch := nextPageRe.FindAllSubmatch(contents, -1)

	for _, m := range nextPageMatch {
		results.Requests = append(results.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: City,
		})
	}

	return results
}
