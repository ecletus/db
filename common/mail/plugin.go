package mail

import (
	admin2 "github.com/aghape/admin"
	"github.com/aghape/admin/adminplugin"
	"github.com/aghape/db"
)

type Plugin struct {
	db.DisDBNames
	adminplugin.AdminNames
}

func (p *Plugin) OnRegister() {
	p.AdminNames.OnInitResources(p, func(e *adminplugin.AdminEvent) {
		e.Admin.AddResource(&QorMail{}, &admin2.Config{Setup: PrepareMailResource})
	})
	p.DBOnMigrateGorm(func(e *db.GormDBEvent) error {
		return e.DB.AutoMigrate(&QorMail{}).Error
	})
}
