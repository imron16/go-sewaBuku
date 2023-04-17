package buku

import "github.com/imron16/go-sewaBuku/entity"

func (s *service) GetBukuById(id int64) (entity.Buku, error) {
	var out *entity.Buku
	result, err := s.bukuRepo.FindBukuById(id)

	if err != nil {
		return *out, err
	}
	out = &result
	return *out, err
}
