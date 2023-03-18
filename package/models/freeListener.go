package models

import "strconv"

type FreeListenerCore struct {
	UserCore
}

type FreeListenerDB struct {
	UserDB
}

type FreeListenerHttp struct {
	UserHTTP `json:"userHttp"`
}

func (em *FreeListenerDB) ToCore() *FreeListenerCore {
	return &FreeListenerCore{
		UserCore: em.UserDB.ToCore(),
	}
}

func (em *FreeListenerDB) FromCore(freeListener *FreeListenerCore) {
	id, _ := strconv.ParseUint(freeListener.Id, 10, 64)
	em.ID = uint(id)
	em.UserDB.FromCore(&freeListener.UserCore)
}

func (ht *FreeListenerHttp) ToCore() *FreeListenerCore {
	return &FreeListenerCore{
		UserCore: ht.UserHTTP.ToCore(),
	}
}

func (ht *FreeListenerHttp) FromCore(freeLister *FreeListenerCore) {
	ht.UserHTTP.FromCore(&freeLister.UserCore)
}
