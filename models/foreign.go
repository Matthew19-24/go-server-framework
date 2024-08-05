package models

import "github.com/google/uuid"

type Foreign struct {
	ForeignUUID uuid.UUID `json:"foreignUUID" gorm:"primary_key;column:foreign_uuid;type:varchar(36)"`
	ForeignName string `json:"foreignName" gorm:"column:foreignName;type:varchar(256)"`
	ForeignExample uuid.UUID `json:"foreignExample" gorm:"column:foreign_example;varchar(36)"`
	Example Example `json:"-" gorm:"foreignkey:foreign_example"`
}
