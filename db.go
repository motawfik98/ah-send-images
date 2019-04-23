package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mssql"
  "github.com/speps/go-hashids"
  "golang.org/x/crypto/bcrypt"
)

func initDB() {
  db, err := gorm.Open("mssql", "sqlserver://remote:mohamed@localhost:1433?database=ah_images_go")
  if err == nil {
    db.AutoMigrate(&MainDepartment{}, &SecondaryDepartment{}, &User{})
  }
}

func generateHash(ID int) string {
  hd := hashids.NewData()
  hd.Salt = "xOBtdmJZxRcz^jkkyHfkrkT1*02bJUn+YQts0*xCeka%cGHCN1fjaC*faFtY"
  hd.MinLength = 8
  h, _ := hashids.NewWithData(hd)
  e, _ := h.Encode([]int{ID})
  return e
}
