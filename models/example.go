package models

import "github.com/google/uuid"

type Example struct {
	ExampleUUID uuid.UUID `json:"exampleUUID" gorm:"primary_key;column:example_uuid;type:varchar(36)"`
	ExampleName string `json:"exampleName" gorm:"column:exampleName;type:varchar(256)"` 
}
