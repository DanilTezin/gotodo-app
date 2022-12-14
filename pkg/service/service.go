package service

import "github.com/DanilTezin/go-todo/pkg/repository"

type Authorization interface {
}
type TodoList interface {
}
type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
