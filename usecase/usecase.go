package usecase

import "github.com/yakan15/clean-transaction/store"

type UseCase interface {
	Test1()
	Test2()
}

type useCase struct {
	s store.Store
}

func NewUseCase(s store.Store) *useCase {
	return &useCase{s: s}
}

func (u *useCase) Test1() {

}
