package model

import (
	"time"

	"gorm.io/gorm"
)

type TypeMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}

type ObjectMeta struct {
	ID         uint64         `json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id"`
	InstanceID string         `json:"instanceID,omitempty" gorm:"unique;column:instanceID;type:varchar(48);not null"`
	Name       string         `json:"name,omitempty" gorm:"column:name;type:varchar(64);not null" validate:"name"`
	CreatedAt  time.Time      `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deletedAt"`
}

type ListMeta struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}

type GetOptions struct {
	TypeMeta `json:",inline"`
}

type DeleteOptions struct {
	TypeMeta `json:",inline"`
	Unscoped bool `json:"unscoped"`
}

type CreateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
}

type UpdateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
}

type ListOptions struct {
	TypeMeta `json:",inline"`

	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`
	Offset         *int64 `json:"offset,omitempty" form:"offset"`
	Limit          *int64 `json:"limit,omitempty" form:"limit"`
}
