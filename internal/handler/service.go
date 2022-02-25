package handler

import "diplom/internal"

type Service struct {
	resultS internal.ResultSetT
	//result  internal.ResultT
}

//
func NewService(result internal.ResultSetT) *Service {
	return &Service{
		resultS: result,
	}
}
