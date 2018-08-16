package callbacks

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"gopkg.in/mgo.v2/bson"
)

var GEN_SERIAL_CALLBACK = PREFIX + ".gen_serial"

func IsGeneratedIdString(scope *gorm.Scope) bool  {
	value := scope.Value
	reflectValue := reflect.ValueOf(value)

	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	pf := scope.PrimaryField()
	return pf != nil && pf.Field.Kind() == reflect.String && pf.Struct.Tag.Get("serial") == "yes"
}

func genSerial(scope *gorm.Scope) {
	value := scope.Value
	reflectValue := reflect.ValueOf(value)

	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	pf := scope.PrimaryField()

	if pf != nil && pf.Field.Kind() == reflect.String && pf.Struct.Tag.Get("serial") == "yes" {
		if pf.Field.String() == "" {
			v := bson.NewObjectId().String()
			v = v[13:len(v)-2]
			pf.Set(v)
		}
	}
}