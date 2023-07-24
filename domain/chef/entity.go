package chef

import "sync"

type Chef struct {
	ID int `json:"id"`
	sync.RWMutex
}
