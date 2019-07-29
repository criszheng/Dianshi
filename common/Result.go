package common

type Result struct {
	Code    int
	Message string
	Data    interface{}
}

const SUCCESS = 200
const FAIL = 400
