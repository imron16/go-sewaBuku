package buku

import (
	"errors"
	"strings"

	"github.com/imron16/go-sewaBuku/shared/constant"
)

func (s *service) UpdateBuku(judulBuku, kategoriBuku, kodeBuku string, id, jumlahBuku int64) error {
	result, err := s.bukuRepo.FindBukuById(id)
	if err != nil {
		return errors.New(constant.InternalServerError)
	}
	if len(strings.TrimSpace(result.JudulBuku)) == 0 {
		return errors.New(constant.BookTitleAlreadyExist)
	}
	result.JudulBuku = judulBuku
	result.KategoriBuku = kategoriBuku
	result.KodeBuku = kodeBuku
	result.JumlahBuku = jumlahBuku

	err = s.bukuRepo.SaveOrUpdateBuku(result)
	if err != nil {
		return err
	}
	return nil
}
