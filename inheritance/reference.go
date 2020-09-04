package inheritance

import (
	"database/sql"

	"github.com/moisespsena-go/aorm"
)

var EXTRA_COLUMNS_KEY = PREFIX

type ParentModelInterface interface {
	GetQorChild() *Child
	SetQorChild(child *Child)
}

type Child struct {
	Index int
	ID    aorm.ID
}

type InheritedModel struct {
	aorm.ExtraSelectModel
	QorChild *Child `sql:"-"`
}

func (es *InheritedModel) SetAormExtraScannedValues(result map[string]*aorm.ExtraResult) {
	panic("TODO: implement field struct. see above")
	es.ExtraSelectModel.SetAormExtraScannedValues(result)
	if v, ok := es.ExtraScannedValues[EXTRA_COLUMNS_KEY]; ok {
		for _, r := range v.Values {
			if ref, ok := r.(sql.NullString); ok && ref.Valid {
				//es.QorChild = &Child{i, aorm.StrID(ref.String)}
				return
			}
		}
	}
	es.QorChild = nil
}

func (r *InheritedModel) GetQorChild() *Child {
	return r.QorChild
}

func (r *InheritedModel) SetQorChild(child *Child) {
	r.QorChild = child
}
