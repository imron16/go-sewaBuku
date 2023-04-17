package sewa

import "github.com/imron16/go-sewaBuku/entity"

func (s *service) GetSewaById(id int64) (entity.Sewa, error) {
	var out *entity.Sewa
	result, err := s.sewaRepo.FindSewaById(id)

	if err != nil {
		return *out, err
	}
	out = &result
	return *out, err
}
