package people

import (
	"github.com/aghape/admin/adminplugin"
	"github.com/aghape/db"
	"github.com/aghape/db/common/address"
	"github.com/aghape/db/common/mail"
	"github.com/aghape/db/common/phone"
	"github.com/aghape/plug"
)

type Plugin struct {
	db.DisDBNames
	adminplugin.AdminNames
}

func (Plugin) After() []interface{} {
	return []interface{}{&address.Plugin{}, &mail.Plugin{}, &phone.Plugin{}}
}

func (p *Plugin) OnRegister(e *plug.Options) {
	p.AdminNames.OnInitResources(p, func(e *adminplugin.AdminEvent) {
		InitResource(e.Admin)
	})

	p.DBOnMigrateGorm(func(e *db.GormDBEvent) error {
		return e.DB.AutoMigrate(&QorPeopleMedia{}, &QorPeople{}, &QorPeoplePhone{}, &QorPeopleAddress{},
			&QorPeopleMail{}).Error
	})
}
