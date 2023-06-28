package storenum

import (
	"context"
	"math/big"

	"github.com/matrosovm/storenum/internal/pkg/helper"
	pbStorenum "github.com/matrosovm/storenum/pkg/api/storenum"
)

// SetNumber ...
func (s *Service) SetNumber(
	ctx context.Context,
	req *pbStorenum.SetNumberRequest,
) (*pbStorenum.SetNumberResponse, error) {
	number := new(big.Int)
	number, ok := number.SetString(req.GetNumber(), 10)
	if !ok {
		return &pbStorenum.SetNumberResponse{}, &helper.ErrorFormat{}
	}

	s.Lock()
	defer s.Unlock()

	return &pbStorenum.SetNumberResponse{}, s.store.SetNumber(number)
}
