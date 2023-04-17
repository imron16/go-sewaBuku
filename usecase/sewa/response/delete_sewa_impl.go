package sewa

func (s *service) DeleteSewa(id int64) error {
	err := s.sewaRepo.DeleteSewa(id)

	if err != nil {
		return err
	}
	return nil
}
