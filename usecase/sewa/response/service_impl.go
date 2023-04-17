package sewa

import "github.com/imron16/go-sewaBuku/domain/sewa"

type service struct {
	sewaRepo sewa.Repository
}

func NewSewaService(
	sewaRepo sewa.Repository,
) Service {
	return &service{sewaRepo: sewaRepo}
}
