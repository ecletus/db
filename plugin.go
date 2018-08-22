package db

import (
	"github.com/aghape/core"
	"github.com/aghape/db/callbacks"
	"github.com/aghape/plug"
	"github.com/moisespsena-go/aorm"
	"github.com/moisespsena/go-default-logger"
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
		dn.Names = []string{core.DB_SYSTEM}
	}
	return dn.Names
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
	DB *core.DB
}

type GormDBEvent struct {
	plug.PluginEventInterface
	DB *aorm.DB
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
