package credintialservice

import (
	"errors"
	"passwordbot/internal/domain/dto"
	"passwordbot/internal/lib/crypto"
)

func (s *CredintialService) ById(credId uint, key string) (dto.Credintial, error) {
	var res dto.Credintial
    cred, err := s.repo.ById(credId)
	if err != nil {
		return dto.Credintial{}, err
	}

    isOK := crypto.CheckHash(key, cred.KeyHash)
    if !isOK {
        return dto.Credintial{}, errors.New("Неверный ключ")
    }

	login, err := crypto.Decrypt(cred.LoginEncrypted, cred.LoginNonce, key)
    if err != nil {
        return dto.Credintial{}, err
    }

    password, err := crypto.Decrypt(cred.PasswordEncrypted, cred.PasswordNonce, key)
    if err != nil {
        return dto.Credintial{}, err
    }

    res.Name = cred.Name
    res.Description = cred.Description
    res.Login = login
    res.Password = password
    res.Type = cred.Type

    return res, nil
}
