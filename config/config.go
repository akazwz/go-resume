package config

type Conf struct {
	DataBase DataBase `yaml:"database"`
	Server
}
