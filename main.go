package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Test ->
type Test struct {
	ID   BinaryUUID `gorm:"type:binary(16);primary_key;default:(UUID_TO_BIN(UUID()));" json:"id"`
	Name string     `json:"name"`
}

// BeforeCreate -> run before creating the model
func (t *Test) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = BinaryUUID(id)
	return err
}

func main() {
	db, err := NewConnection()
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&Test{})

	createdData := createNewData(db)

	log.Printf("Created Data: %+v \n", createdData)

	data := `{"id": "%s"}`
	data = fmt.Sprintf(data, createdData.ID)

	marshalData := Test{}
	err = json.Unmarshal([]byte(data), &marshalData)
	db.Where("id = ?", marshalData.ID).First(&marshalData)

	log.Printf("Where query on data: %+v \n", marshalData)

}

func createNewData(db *gorm.DB) Test {
	data := Test{Name: "New Data Goes Here!!"}
	db.Create(&data)
	return data
}
