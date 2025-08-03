package credintialservice

import (
	"passwordbot/internal/domain/dto"
	"passwordbot/internal/domain/models"
	"passwordbot/internal/lib/crypto"
	"passwordbot/pkg/logger/sl"
)

func (s *CredintialService) Create(input dto.CredintialInput) error {
	log := s.log.With(sl.Op(op, "Create"))

	var cred models.Credintial
	loginEncryption, loginNonce, err := crypto.Encrypt(input.Login, input.Key)
	if err != nil {
		log.Error(ErrCreateCred.Error(), sl.Err(err))
		return ErrCreateCred
	}

	cred.LoginEncrypted = loginEncryption
	cred.LoginNonce = loginNonce

	passwordEncryption, passwordNonce, err := crypto.Encrypt(input.Password, input.Key)
	if err != nil {
		log.Error(ErrCreateCred.Error(), sl.Err(err))
		return ErrCreateCred
	}

	cred.PasswordEncrypted = passwordEncryption
	cred.PasswordNonce = passwordNonce

	cred.KeyHash = crypto.Hash(input.Key)
	cred.TypeId = input.TypeId
	cred.UserId = input.UserId
	cred.Description = input.Description
	cred.Name = input.Name

	_, err = s.repo.Create(cred)
	if err != nil {
		log.Error(ErrCreateCred.Error(), sl.Err(err))
		return ErrCreateCred
	}

	return nil
}
