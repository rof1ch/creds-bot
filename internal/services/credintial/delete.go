package credintialservice

func (s *CredintialService) Delete(credId uint) error {
    return s.repo.Delete(credId)
}