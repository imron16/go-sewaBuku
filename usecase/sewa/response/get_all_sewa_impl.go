package sewa

import "github.com/imron16/go-sewaBuku/entity"

func (s *service) GetAllSewa() ([]entity.Sewa, error) {
	var out []entity.Sewa
	result, err := s.sewaRepo.All()

	if err != nil {
		return out, err
	}
	out = result
	return out, err
}
