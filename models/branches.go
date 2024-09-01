package models

import (

  "database/sql"
)
type Address struct {}


type Branch struct {
  CodeName sql.NullString `db:"codename"`
  AddressID int `db:"address_id"`
  Address Address // Address struct 
  FreezerSpace int `db:"freezer_space"` 
}

