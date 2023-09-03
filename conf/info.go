package conf

var Conf Config

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Udp      UdpConfig      `yaml:"udp"`
}

type ServerConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type UdpConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
