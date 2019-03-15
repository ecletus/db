package db

import (
	"reflect"

	"github.com/ecletus/plug"
	"github.com/moisespsena/go-edis"
	"github.com/moisespsena/go-error-wrap"
)

type events struct {
	*DBNames
	edis.EventDispatcherInterface
}

func Events(args ...interface{}) (d *events) {
	d = &events{}
	if len(args) == 1 {
		value := reflect.Indirect(reflect.ValueOf(args[0]))
		if f := value.FieldByName("DBNames"); f.IsValid() {
			if f.Kind() == reflect.Struct {
				f = f.Addr()
			}
			d.DBNames = f.Interface().(*DBNames)
		}
		if f := value.FieldByName("EventDispatcher"); f.IsValid() {
			if f.Kind() == reflect.Struct {
				f = f.Addr()
			}
			d.EventDispatcherInterface = f.Interface().(edis.EventDispatcherInterface)
		}
		if f := value.FieldByName("EventDispatcherInterface"); f.IsValid() {
			if f.Kind() == reflect.Struct {
				f = f.Addr()
			}
			d.EventDispatcherInterface = f.Interface().(edis.EventDispatcherInterface)
		}
		return
	}
	for _, arg := range args {
		switch at := arg.(type) {
		case *DBNames:
			d.DBNames = at
		case edis.EventDispatcherInterface:
			d.EventDispatcherInterface = at
		}
	}
	return
}

func (dn *events) DBEachApllyE(en func(name string) string, cb func(e *DBEvent) error) *events {
	for _, dbName := range dn.GetNames() {
		dn.On(en(dbName), func(e plug.EventInterface) error {
			err := cb(e.(*DBEvent))
			return errwrap.Wrap(err, e.Name())
		})
	}
	return dn
}

func (dn *events) DBEachAplly(en func(name string) string, cb func(e *DBEvent)) *events {
	return dn.DBEachApllyE(en, func(e *DBEvent) error {
		cb(e)
		return nil
	})
}

func (dn *events) DBEachGormApllyE(en func(name string) string, cb func(e *DBEvent) error) *events {
	for _, dbName := range dn.GetNames() {
		dn.On(en(dbName), func(e plug.EventInterface) error {
			err := cb(e.(*DBEvent))
			return errwrap.Wrap(err, e.Name())
		})
	}
	return dn
}

func (dn *events) DBEachRawApllyE(en func(name string) string, cb func(e *DBEvent) error) *events {
	for _, dbName := range dn.GetNames() {
		dn.On(en(dbName), func(e plug.EventInterface) error {
			err := cb(e.(*DBEvent))
			return errwrap.Wrap(err, e.Name())
		})
	}
	return dn
}

func (dn *events) DBEachGormAplly(en func(name string) string, cb func(e *DBEvent)) *events {
	return dn.DBEachGormApllyE(en, func(e *DBEvent) error {
		cb(e)
		return nil
	})
}

func (dn *events) DBOnInitE(cb func(e *DBEvent) error) *events {
	return dn.DBEachApllyE(EInit, cb)
}

func (dn *events) DBOnInit(cb func(e *DBEvent)) *events {
	return dn.DBEachAplly(EInit, cb)
}

func (dn *events) DBOnInitGormE(cb func(e *DBEvent) error) *events {
	return dn.DBEachGormApllyE(EInitGorm, cb)
}

func (dn *events) DBOnInitGorm(cb func(e *DBEvent)) *events {
	return dn.DBEachGormAplly(EInitGorm, cb)
}

func (dn *events) DBOnMigrate(cb func(e *DBEvent) error) *events {
	return dn.DBEachApllyE(EMigrate, cb)
}