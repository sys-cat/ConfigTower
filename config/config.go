package Config

import (
  "github.com/BurntSushi/toml"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
)

type Config struct {
  Serv Server
  DB MySQL
}

type Server struct {
  PORT string `toml:"PORT"`
}

type MySQL struct {
  HOST string `toml:"HOST"`
  PORT string `toml:"PORT"`
  USER string `toml:"USER"`
  PASS string `toml:"PASS"`
}

func main() {}
