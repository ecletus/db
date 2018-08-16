package geocode

import (
	"fmt"

	"github.com/aghape/aghape"
	"github.com/aghape/aghape/helpers"
)

func Migrate(DB *qor.DB) error {
	db, rawDB := DB.DB, DB.Raw
	country := db.NewScope(&GeoCodeCountry{})
	state := db.NewScope(&GeoCodeRegion{})
	key, err := helpers.CheckReturnError(
		func() (key string, err error) {
			return "MigrateModels", db.AutoMigrate(&GeoCodeCdhCountryCode{}, &GeoCodeCdhStateCode{}, &GeoCodeCountry{},
				&GeoCodeRegion{}).Error
		},
		func() (key string, err error) {
			return country.TableName() + ".CreateIndex.Code2", country.DB().AddIndex("code2", "code2").Error
		},
		func() (key string, err error) {
			return country.TableName() + ".CreateIndex.Code3", country.DB().AddIndex("code3", "code3").Error
		},
		func() (key string, err error) {
			return state.TableName() + ".CreateIndex.CountryID", state.DB().AddIndex("country_id", "country_id").Error
		},
		func() (key string, err error) {
			return state.TableName() + ".CreateIndex.Level", state.DB().AddIndex("level", "level").Error
		},
		func() (key string, err error) {
			var v int
			err = db.Model(country).Count(&v).Error
			if err != nil {
				return country.TableName() + ".Count", err
			}
			if v == 0 {
				return ImportPGSQLData(rawDB)
			}
			return
		},
	)
	if err != nil {
		return fmt.Errorf("qor/db.common.geocode.setup.MigrateDB.%v: failed to migrate DB: %v", key, err)
	}
	return nil
}
