package models

type Response[T any] struct {
	Code   int
	Status string
	Result T
}

type Response_arr[T any] struct {
	Code   int
	Status string
	Result []T
}
