package go_micro_srv_user

import (
	"log"

	"github.com/jinzhu/gorm"

	uuid "github.com/nu7hatch/gouuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	u4, err := uuid.NewV4()
	if err != nil {
		log.Printf("Error creating UUID: %v", err)
		return err
	}
	return scope.SetColumn("Id", u4.String())
}
