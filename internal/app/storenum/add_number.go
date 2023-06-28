package storenum

import (
	"context"
	"fmt"
	"math/big"

	"github.com/matrosovm/storenum/internal/pkg/helper"
	pbStorenum "github.com/matrosovm/storenum/pkg/api/storenum"
)

// AddNumber ...
func (s *Service) AddNumber(
	ctx context.Context,
	req *pbStorenum.AddNumberRequest,
) (*pbStorenum.AddNumberResponse, error) {
	number := new(big.Int)
	number, ok := number.SetString(req.GetNumber(), 10)
	if !ok {
		return &pbStorenum.AddNumberResponse{}, &helper.ErrorFormat{}
	}

	if number.Cmp(big.NewInt(0)) != 1 {
		return &pbStorenum.AddNumberResponse{}, fmt.Errorf("number must be natural")
	}

	s.Lock()
	defer s.Unlock()

	storedNumber, err := s.store.GetNumber()
	if err != nil {
		return &pbStorenum.AddNumberResponse{}, err
	}

	storedNumber = storedNumber.Add(storedNumber, number)
	err = s.store.SetNumber(storedNumber)
	if err != nil {
		return &pbStorenum.AddNumberResponse{}, err
	}

	return &pbStorenum.AddNumberResponse{
		Number: storedNumber.String(),
	}, nil
}
