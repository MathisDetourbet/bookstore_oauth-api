package db

import (
	"github.com/MathisDetourbet/bookstore_oauth-api/src/domain/access_token"
	"github.com/MathisDetourbet/bookstore_users-api/utils/errors"
)

func New() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}
