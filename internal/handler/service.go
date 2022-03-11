package handler

import "diplom/internal"

type Service struct {
	resultS internal.ResultSetT
	result  internal.ResultT
}

//
func NewService(result internal.ResultT) *Service {
	return &Service{
		result: result,
	}
}
