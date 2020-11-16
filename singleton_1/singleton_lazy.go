package singleton

var lazySingleton *Singleton

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		lazySingleton = &Singleton{}
	}
	return lazySingleton
}
