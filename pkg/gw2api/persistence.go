package gw2api

import (
	"time"
)

type Gw2Model struct {
	DbCreated time.Time
	DbUpdated time.Time
}

func (ent *Gw2Model) BeforeSave() (err error) {
	ent.DbUpdated = time.Now().UTC()
	return nil
}

func (ent *Gw2Model) BeforeCreate() (err error) {
	ent.DbCreated = time.Now().UTC()
	return nil
}
