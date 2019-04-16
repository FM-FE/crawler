package engine

type Request struct {
	Url       string
	ParserFunc func([]byte) ParserRequests
}

type ParserRequests struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParserRequests {
	return ParserRequests{}
}
