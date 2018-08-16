package address

import (
	"github.com/aghape/admin/adminplugin"
	"github.com/aghape/db"
	"github.com/aghape/db/common/geocode"
	"github.com/aghape/db/common/phone"
)

type Plugin struct {
	db.DisDBNames
	adminplugin.AdminNames
}

func (Plugin) After() []interface{} {
	return []interface{}{&geocode.Plugin{}, &phone.Plugin{}}
}

func (p *Plugin) OnRegister() {
	p.AdminNames.OnInitResources(p, func(e *adminplugin.AdminEvent) {
		e.Admin.AddResource(&QorAddress{}, &adminplugin.Config{Setup: PrepareResource})
	})
	p.DBOnMigrateGorm(func(e *db.GormDBEvent) error {
		return e.DB.AutoMigrate(&QorAddress{}, &QorAddressPhone{}).Error
	})
}
