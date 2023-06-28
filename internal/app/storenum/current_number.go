package storenum

import (
	"context"

	pbStorenum "github.com/matrosovm/storenum/pkg/api/storenum"
)

// CurrentNumber ...
func (s *Service) CurrentNumber(
	ctx context.Context,
	req *pbStorenum.CurrentNumberRequest,
) (*pbStorenum.CurrentNumberResponse, error) {
	s.RLock()
	defer s.RUnlock()
	
	number, err := s.store.GetNumber()
	if err != nil {
		return &pbStorenum.CurrentNumberResponse{}, err
	}

	return &pbStorenum.CurrentNumberResponse{
		Number: number.String(),
	}, nil
}
