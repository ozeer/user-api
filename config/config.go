package config

type Conf struct {
	App    SectionApp    `mapstructure:"app"`
	Server SectionServer `mapstructure:"server"`
	Mysql  SectionMysql  `mapstructure:"mysql"`
}

type SectionApp struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

type SectionServer struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type SectionMysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       string `mapstructure:"db"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
