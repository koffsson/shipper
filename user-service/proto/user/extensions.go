package go_micro_srv_user

import (
	"log"

	"github.com/jinzhu/gorm"
	// "github.com/satori/go.uuid"

	uuid "github.com/nu7hatch/gouuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	// uuid, err := uuid.NewV4()
	// if err != nil {
	// 	log.Printf("Something went wrong: %v", err)
	// 	return err
	// }
	// return scope.SetColumn("Id", uuid.String())

	// u1 := uuid.Must(uuid.NewV4())
	// scope.SetColumn("Id", u1.String())
	// return nil

	u4, err := uuid.NewV4()
	if err != nil {
		log.Printf("Error creating UUID: %v", err)
		return err
	}
	return scope.SetColumn("Id", u4.String())
}
