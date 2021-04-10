## Gorm Binary UUID `binary(16)`
Integrating GORM uuid implementation using new `binary(16)` uuid implementation available in `mysql version 8`.

### `BinaryUUID`
- Wraps `uuid.UUID` inside and converts into binary when needed.
- Implements Gorm `Scan` and `Value` for data type conversion.
- Implements `GormDataType` for data type in sql.
- `MarshalJSON` and `UnmarshalJSON` implementation for json encode/decode.


##### GORM's `BeforeCreate`
`BeforeCreate` hook for creating new uuids whenever new data has to be inserted.

```go
// BeforeCreate ->
func (t *Test) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = BinaryUUID(id)
	return err
}
```


### Usage

##### Where statement:
```go
model := Test{}
db.Find(&model, "id = ?", ParseUUID("ed67a4b2-8f77-4d18-8c58-0508e7b207e8"))
log.Printf("data: %+v\n", model)
```

##### Create statement: 
Will autogenerate uuid automatically because of `BeforeCreate` hook added
```go
data := Test{Name: fmt.Sprintf("Time is: %s", time.Now())}
db.Create(&data)
```

**Check [this article](https://dipesh-kc.medium.com/implementation-of-uuid-and-binary-16-in-gorm-v2-1c329c352c91?sk=04dd6a121675ed4f2a635eabc3c3592b) for more information**
