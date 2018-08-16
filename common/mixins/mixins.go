package mixins

import (
	"time"
)

type CreationMixin struct {
	CreatedAt *time.Time
	CreatedById string `gorm:"size:255"`
}

type UpdationMixin struct {
	UpdatedAt *time.Time
	UpdatedById string `gorm:"size:24"`
}

type DeletionMixin struct {
	DeletedAt *time.Time
	DeletedById string `gorm:"size:24"`
}

type CreationUpdationMixin struct {
	CreationMixin
	UpdationMixin
}

type RecordInfoMixin struct {
	CreationMixin
	UpdationMixin
	DeletionMixin
}
