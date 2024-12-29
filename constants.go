package starter

type ParserValueType int

const (
	Num ParserValueType = iota
	Text
)

var parserValueTypeMap = map[ParserValueType]string{
	Num:  "Num",
	Text: "Text",
}

func (p ParserValueType) String() string {
	return parserValueTypeMap[p]
}

type ResponseStatus int

const (
	ResponseSuccess ResponseStatus = iota
	ResponseError
	ResponseFail
)

var responseStatusString = map[ResponseStatus]string{
	ResponseSuccess: "success",
	ResponseError:   "error",
	ResponseFail:    "fail",
}

func (r ResponseStatus) String() string {
	return responseStatusString[r]
}
