package callbacks

import (
	"reflect"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/aghape/db/common/utils"
	"github.com/aghape/common"
)

var SET_UPDATION_INFO_CALLBACK = PREFIX + ".set_updation_info"
var SET_CREATION_INFO_CALLBACK = PREFIX + ".set_updation_info"
var SET_DELETION_INFO_CALLBACK = PREFIX + ".set_deletion_info"

func setUpdateInfo(user common.User, obj reflect.Value) {
	userID := user.GetID()
	utils.SetStringValueIfEmpty(obj, "UpdatedById", &userID)
	now := time.Now()
	utils.SetTimeValueIfEmpty(obj, "UpdatedAt", &now)
	utils.SetTimeValue(obj, "DeletedAt", nil, func(v *time.Time) bool {
		if (v == nil || v.IsZero()) {
			utils.SetStringValue(obj, "DeletedById", nil, func(v *string) bool {
				return true
			})
		}
		return true
	})
}

func setUpdateInfoCallback(scope *gorm.Scope) {
	if user := common.GetUserFromScope(scope); user != nil {
		if scope.Value == nil {
			return
		}

		if indirectScopeValue := scope.IndirectValue(); indirectScopeValue.Kind() == reflect.Slice {
			if indirectScopeValue.Len() > 0 {
				for i := 0; i < indirectScopeValue.Len(); i++ {
					setUpdateInfo(user, indirectScopeValue.Index(i))
				}
			}
		} else {
			setUpdateInfo(user, indirectScopeValue)
		}
	}
}

func setCreateInfo(user common.User, obj reflect.Value) {
	userID := user.GetID()
	utils.SetStringValueIfEmpty(obj, "CreatedById", &userID)
	now := time.Now()
	utils.SetTimeValueIfEmpty(obj, "CreatedAt", &now)
}

func setCreateInfoCallback(scope *gorm.Scope) {
	if user := common.GetUserFromScope(scope); user != nil {
		if scope.Value == nil {
			return
		}

		if indirectScopeValue := scope.IndirectValue(); indirectScopeValue.Kind() == reflect.Slice {
			if indirectScopeValue.Len() > 0 {
				for i := 0; i < indirectScopeValue.Len(); i++ {
					setCreateInfo(user, indirectScopeValue.Index(i))
				}
			}
		} else {
			setCreateInfo(user, indirectScopeValue)
		}
	}
}

func setDeleteInfo(user common.User, obj reflect.Value) {
	userID := user.GetID()
	utils.SetStringValueIfEmpty(obj, "DeletedById", &userID)
	now := time.Now()
	utils.SetTimeValueIfEmpty(obj, "DeletedAt", &now)
}

func setDeleteInfoCallback(scope *gorm.Scope) {
	if user := common.GetUserFromScope(scope); user != nil {
		if scope.Value == nil {
			return
		}

		if indirectScopeValue := scope.IndirectValue(); indirectScopeValue.Kind() == reflect.Slice {
			if indirectScopeValue.Len() > 0 {
				for i := 0; i < indirectScopeValue.Len(); i++ {
					setDeleteInfo(user, indirectScopeValue.Index(i))
				}
			}
		} else {
			setDeleteInfo(user, indirectScopeValue)
		}
	}
}
