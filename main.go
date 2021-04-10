package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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

	// where statement by unmarshaling string to binary uuid
	data := `{"id": "%s"}`
	data = fmt.Sprintf(data, createdData.ID)
	marshalData := Test{}
	err = json.Unmarshal([]byte(data), &marshalData)
	db.Where("id = ?", marshalData.ID).First(&marshalData)
	log.Printf("Where query on data: %+v \n", marshalData)

	// Where statement
	model := Test{}
	db.Find(&model, "id = ?", ParseUUID("ed67a4b2-8f77-4d18-8c58-0508e7b207e8"))
	log.Printf("data: %+v\n", model)
}

func createNewData(db *gorm.DB) Test {
	data := Test{Name: fmt.Sprintf("Time is: %s", time.Now())}
	db.Create(&data)
	return data
}
