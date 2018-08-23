package callbacks

import (
	"github.com/aghape/validations"
	"github.com/moisespsena-go/aorm"
)

// Register register callback into GORM DB
func Register(db *aorm.DB) *aorm.DB {
	db = validations.RegisterCallbacks(db)
	db.Callback().Create().Before(validations.VALIDATE_CALLBACK).Register(CLEAN_CALLBACK, cleanCallback)

	db.Callback().Update().Before(validations.VALIDATE_CALLBACK).Register(CLEAN_CALLBACK, cleanCallback)
	return db
}
