package main

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  "../config"
)

type User struct {
  ID int64
  NAME string
  DESC string
  PASS string
  MAIL string
  AUTH string
}

func Add(seed User)result bool {}

func Edit(name string, desc string, pass string, mail string)result bool {}

func Delete(id int64)result bool{}

func BuildAuthKey(pass_key string)auth string {}
