package entity

// Online 上线的结构体，主要包含id、名字、ip地址
type Online struct {
	Id   int64  `yaml:"id"`
	Name string `yaml:"name"`
	Ip   string `yaml:"ip"`
}
