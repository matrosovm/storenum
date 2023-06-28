package storenum

import (
	"sync"

	"github.com/matrosovm/storenum/internal/pkg/database"
	pbStorenum "github.com/matrosovm/storenum/pkg/api/storenum"
)

// Service ...
type Service struct {
	pbStorenum.UnimplementedStorenumServer
	sync.RWMutex
	
	store database.Store
}

// NewService ...
func NewService(store database.Store) pbStorenum.StorenumServer {
	return &Service{
		store: store,
	}
}
