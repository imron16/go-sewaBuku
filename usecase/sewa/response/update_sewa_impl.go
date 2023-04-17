package sewa

import (
	"errors"
	"strings"

	"github.com/imron16/go-sewaBuku/shared/constant"
)

func (s *service) UpdateSewa(anggota, judulBuku string, id int64) error {
	result, err := s.sewaRepo.FindSewaById(id)

	if err != nil {
		return errors.New(constant.InternalServerError)
	}
	if len(strings.TrimSpace(result.JudulBuku)) == 0 {
		return errors.New(constant.BookTitleAlreadyExist)
	}

	result.Anggota = anggota
	result.JudulBuku = judulBuku

	err = s.sewaRepo.SaveOrUpdateSewa(result)

	if err != nil {
		return err
	}
	return nil
}
