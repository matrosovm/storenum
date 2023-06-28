package database

import (
	"math/big"
	"os"

	"github.com/matrosovm/storenum/internal/pkg/helper"
)

// Store ...
type Store interface {
	SetNumber(*big.Int) error
	GetNumber() (*big.Int, error)
}

type storeImpl struct {
	fileName string
}

// NewStore ...
func NewStore(fileName string) Store {
	return &storeImpl{
		fileName: fileName,
	}
}

func (s *storeImpl) SetNumber(number *big.Int) error {
	return os.WriteFile(s.fileName, []byte(number.String()), 0222)
}

func (s *storeImpl) GetNumber() (*big.Int, error) {
	buff, err := os.ReadFile(s.fileName)
	if err != nil {
		return nil, err
	}

	number := new(big.Int)
	number, ok := number.SetString(string(buff), 10)
	if !ok {
		return nil, &helper.ErrorFormat{}
	}

	return number, nil
}
