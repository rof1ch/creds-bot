package typeservice

func (s *TypeService) Delete(typeId uint) error {
	return s.repo.Delete(typeId)
}
