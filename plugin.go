package db

import (
	"github.com/jinzhu/gorm"
	"github.com/moisespsena/go-default-logger"
	"github.com/moisespsena/go-error-wrap"
	"github.com/aghape/db/callbacks"
	"github.com/aghape/plug"
	"github.com/aghape/aghape"
)

var (
	E_INIT         = PREFIX + ":init"
	E_INIT_DB      = E_INIT + ":db"
	E_INIT_GORM    = E_INIT + ":gorm"
	E_MIGRATE      = PREFIX + ":migrate"
	E_MIGRATE_DB   = E_MIGRATE + ":db"
	E_MIGRATE_GORM = E_MIGRATE + ":gorm"
	log            = defaultlogger.NewLogger(PREFIX)
)

type DBNames struct {
	Names []string
}

func (dn *DBNames) GetNames() []string {
	if len(dn.Names) == 0 {
		dn.Names = []string{qor.DB_SYSTEM}
	}
	return dn.Names
}

type DisDBNames struct {
	DBNames
	plug.EventDispatcher
}

func (dn *DisDBNames) DBEachApllyE(en func(name string) string, cb func(e *DBEvent) error) {
	for _, dbName := range dn.GetNames() {
		dn.On(en(dbName), func(e plug.PluginEventInterface) error {
			err := cb(e.(*DBEvent))
			return errwrap.Wrap(err, e.Name())
		})
	}
}

func (dn *DisDBNames) DBEachAplly(en func(name string) string, cb func(e *DBEvent)) {
	dn.DBEachApllyE(en, func(e *DBEvent) error {
		cb(e)
		return nil
	})
}

func (dn *DisDBNames) DBEachGormApllyE(en func(name string) string, cb func(e *GormDBEvent) error) {
	for _, dbName := range dn.GetNames() {
		dn.On(en(dbName), func(e plug.PluginEventInterface) error {
			err := cb(e.(*GormDBEvent))
			return errwrap.Wrap(err, e.Name())
		})
	}
}

func (dn *DisDBNames) DBEachGormAplly(en func(name string) string, cb func(e *GormDBEvent)) {
	dn.DBEachGormApllyE(en, func(e *GormDBEvent) error {
		cb(e)
		return nil
	})
}

func (dn *DisDBNames) DBOnInitE(cb func(e *DBEvent) error) {
	dn.DBEachApllyE(EInit, cb)
}

func (dn *DisDBNames) DBOnInit(cb func(e *DBEvent)) {
	dn.DBEachAplly(EInit, cb)
}

func (dn *DisDBNames) DBOnInitGormE(cb func(e *GormDBEvent) error) {
	dn.DBEachGormApllyE(EInitGorm, cb)
}

func (dn *DisDBNames) DBOnInitGorm(cb func(e *GormDBEvent)) {
	dn.DBEachGormAplly(EInitGorm, cb)
}

func (dn *DisDBNames) DBOnMigrate(cb func(e *DBEvent) error) {
	dn.DBEachApllyE(EMigrate, cb)
}

func (dn *DisDBNames) DBOnMigrateGorm(cb func(e *GormDBEvent) error) {
	dn.DBEachGormApllyE(EMigrateGorm, cb)
}

func EInit(name string) string {
	if name == "" {
		panic("name is blank")
	}
	return E_INIT_DB + ":" + name
}

func EInitGorm(name string) string {
	if name == "" {
		panic("name is blank")
	}
	return E_INIT_GORM + ":" + name
}

func EMigrate(name string) string {
	if name == "" {
		panic("name is blank")
	}
	return E_MIGRATE_DB + ":" + name
}

func EMigrateGorm(name string) string {
	if name == "" {
		panic("name is blank")
	}
	return E_MIGRATE_GORM + ":" + name
}

type DBEvent struct {
	plug.PluginEventInterface
	DB *qor.DB
}

type GormDBEvent struct {
	plug.PluginEventInterface
	DB *gorm.DB
}

type Plugin struct {
	plug.EventDispatcher
}

func (p *Plugin) init(e plug.PluginEventInterface) (err error) {
	log.Debug("trigger")
	dis := e.PluginDispatcher()
	defer e.With(E_INIT_DB)()
	err = dis.TriggerPlugins(e)
	if err != nil {
		return
	}
	defer e.With(E_INIT_GORM)()
	err = dis.TriggerPlugins(e)
	return
}

func (p *Plugin) migrate(e plug.PluginEventInterface) (err error) {
	log.Debug("trigger")
	dis := e.PluginDispatcher()
	defer e.With(E_MIGRATE_DB)()
	err = dis.TriggerPlugins(e)
	if err != nil {
		return
	}
	defer e.With(E_MIGRATE_GORM)()
	err = dis.TriggerPlugins(e)
	return
}

func (p *Plugin) OnRegister() {
	p.EnableAnyTrigger()
	p.On(plug.EAll(E_INIT_GORM), func(e plug.PluginEventInterface) {
		edb := e.(*GormDBEvent)
		callbacks.Register(edb.DB)
		return
	})

	p.On(E_INIT, p.init)
	p.On(E_MIGRATE, p.migrate)
}
