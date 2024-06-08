package responseDto

type Response struct {
	Status string
	Data   interface{}
}

type Error struct {
	Cause string `json:"cause"`
}

func GetErrorResponseObj(cause string) Response {
	return Response{Status: "Fail", Data: Error{cause}}
}

func GetSuccessResponseObj(data interface{}) Response {
	return Response{Status: "Success", Data: data}
}
