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
	ID    string
}

type InheritedModel struct {
	aorm.ExtraSelectModel
	QorChild *Child `sql:"-"`
}

func (es *InheritedModel) SetGormExtraScannedValues(result map[string]*aorm.ExtraResult) {
	es.ExtraSelectModel.SetGormExtraScannedValues(result)
	if v, ok := es.ExtraScannedValues[EXTRA_COLUMNS_KEY]; ok {
		for i, r := range v.Values {
			if ref, ok := r.(sql.NullString); ok && ref.Valid {
				es.QorChild = &Child{i, ref.String}
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
