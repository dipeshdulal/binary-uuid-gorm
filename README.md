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

