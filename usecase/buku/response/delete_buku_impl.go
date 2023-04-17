package buku

func (s *service) DeleteBuku(id int64) error {
	err := s.bukuRepo.DeleteBuku(id)

	if err != nil {
		return err
	}
	return nil
}
