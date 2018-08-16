package callbacks

import (
	"github.com/jinzhu/gorm"
	"github.com/aghape/validations"
)

// Register register callback into GORM DB
func Register(db *gorm.DB) *gorm.DB {
	db = validations.RegisterCallbacks(db)
	db.Callback().Create().Before(validations.VALIDATE_CALLBACK).Register(CLEAN_CALLBACK, cleanCallback)
	db.Callback().Create().Before(validations.VALIDATE_CALLBACK).Register(SET_CREATION_INFO_CALLBACK, setCreateInfoCallback)
	db.Callback().Create().Before(validations.VALIDATE_CALLBACK).Register(GEN_SERIAL_CALLBACK, genSerial)

	db.Callback().Update().Before(validations.VALIDATE_CALLBACK).Register(CLEAN_CALLBACK, cleanCallback)
	db.Callback().Update().Before(validations.VALIDATE_CALLBACK).Register(SET_UPDATION_INFO_CALLBACK, setUpdateInfoCallback)

	db.Callback().Delete().Before("gorm:delete").Register(SET_DELETION_INFO_CALLBACK, setUpdateInfoCallback)
	return db
}
