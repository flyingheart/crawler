package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

const (
	CityRequest = iota
	ProfileRequest
	ProfileDetails
)

type ParseResult struct {
	Requests []Request
	Items    []interface{}
	ItemType int
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
