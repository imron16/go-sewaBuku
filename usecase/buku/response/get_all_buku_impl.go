package buku

import "github.com/imron16/go-sewaBuku/entity"

func (s *service) GetAllBuku() ([]entity.Buku, error) {
	var out []entity.Buku
	result, err := s.bukuRepo.All()
	if err != nil {
		return out, err
	}
	out = result
	return out, err
}
