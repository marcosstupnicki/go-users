package config

type Database struct {
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}

type Config struct {
	Database Database `yaml:"database"`
}

type Configs struct {
	Scope map[string]Config
}