package callbacks

import (
	"github.com/jinzhu/gorm"
)

var CLEAN_CALLBACK = PREFIX + ".clean"
var CLEAN_SKIP = CLEAN_CALLBACK + "_skip"

func cleanCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		if result, ok := scope.DB().Get(CLEAN_SKIP); !(ok && result.(bool)) {
			if !scope.HasError() {
				scope.CallMethod("Clean")
			}
		}
	}
}
