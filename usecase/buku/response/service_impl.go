package buku

import "github.com/imron16/go-sewaBuku/domain/buku"

type service struct {
	bukuRepo buku.Repository
}

func NewBukuService(
	bukuRepo buku.Repository,
) Service {
	return &service{bukuRepo: bukuRepo}
}
