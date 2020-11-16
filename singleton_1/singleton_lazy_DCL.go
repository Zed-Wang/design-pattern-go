package singleton

import "sync"

var (
	lazySingletonDCL *Singleton
	once             sync.Once
)

// Double Check Lock
func GetLazyDCLInstance() *Singleton {
	once.Do(func() {
		lazySingletonDCL = &Singleton{}
	})
	return lazySingletonDCL
}
