package common

import (
	"gopkg.in/mgo.v2/bson"
)

type IDStringSerial struct {
	ID string `gorm:"size:24;primary_key" serial:"yes"`
}

func (p *IDStringSerial) GetID() string {
	return p.ID
}

func (p *IDStringSerial) BeforeCreate() (err error) {
	p.ID = bson.NewObjectId().Hex()
	return nil
}

type IDString struct {
	ID string `gorm:"size:24;primary_key"`
}

func (p *IDString) GetID() string {
	return p.ID
}

func (p *IDString) SetID(value string) {
	p.ID = value
}

type Model struct {
	IDStringSerial
}

type NamedModel struct {
	Model
	Name string `gorm:"size:255"`
}

type Address struct {
	Model
	Name        string `gorm:"size:255"`
	Country     string `gorm:"size:255"`
	Address1    string `gorm:"size:255"`
	Address2    string `gorm:"size:255"`
	TownCity    string `gorm:"size:255"`
	Postcode    string `gorm:"size:255"`
	Phone       string `gorm:"size:255"`
	ContactName string `gorm:"size:255"`
}

type Phone struct {
	Model
	Name   string
	Number string
}

type Mail struct {
	Model
	Name    string
	Address string
}

type Social struct {
	Model
	Url string
}
