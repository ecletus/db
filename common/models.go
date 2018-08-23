package common

import "github.com/moisespsena-go/aorm"

type (
	IDStringSerial = aorm.KeyStringSerial
	IDString       = aorm.KeyString
)

type Model struct {
	aorm.KeyStringSerial
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
