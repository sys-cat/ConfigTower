package Config

import (
  "github.com/BurntSushi/toml"
)

type Config struct {
  Serv Server
  DB MySQL
  Auth Auth
}

type Server struct {
  PORT string `toml:"PORT"`
  ENV string `toml:"ENV"`
}

type MySQL struct {
  HOST string `toml:"HOST"`
  PORT string `toml:"PORT"`
  USER string `toml:"USER"`
  PASS string `toml:"PASS"`
}

type Auth struct {
  SEC string `toml:"SECRET"`
  SALT string `toml:"SALT"`
}

func Conf()(conf Config) {
  _, err := toml.Decodefile("./base.toml", &conf)
  if err != nil {
    panic(err)
  }
  return conf
}
