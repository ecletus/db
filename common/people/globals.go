package people

import (
	"github.com/jinzhu/gorm"
	"github.com/moisespsena/go-i18n-modular/i18nmod"
	"github.com/moisespsena/go-path-helpers"
	"github.com/moisespsena/template/html/template"
	"github.com/aghape/admin"
	"github.com/aghape/admin/tabs"
	"github.com/aghape/aghape"
	"github.com/aghape/aghape/utils"
)

var DefaultTab = &tabs.Tab{
	Default: true,
	Title:   "All",
	Handler: func(res *admin.Resource, context *qor.Context, db *gorm.DB) *gorm.DB {
		return db
	},
}

var ImageTag, _ = template.New("qor:db:common.people.tag.image").Parse("<img src=\"{{.}}\"></img>")
var PeopleTabs = []*tabs.Tab{
	DefaultTab,
	{"Individual", "", "", func(res *admin.Resource, context *qor.Context, db *gorm.DB) *gorm.DB {
		return db
	}, false},
	{"Business", "", "", func(res *admin.Resource, context *qor.Context, db *gorm.DB) *gorm.DB {
		return db
	}, false},
}

var (
	PKG        = path_helpers.GetCalledDir()
	I18N_GROUP = i18nmod.PkgToGroup(PKG)
)

func init() {
	group := I18N_GROUP + ".QorPeople"

	for _, scope := range PeopleTabs {
		scope.Path = utils.ToParamString(scope.Title)
		scope.TitleKey = group + ".scopes." + scope.Title
	}
}
