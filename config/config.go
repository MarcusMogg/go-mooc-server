package config

type Config struct {
	AesKey string `mapstructure:"asekey" json:"asekey" yaml:"asekey"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Addr   int    `mapstructure:"addr" json:"addr" yaml:"addr"`
}

type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Parm     string `mapstructure:"parm" json:"parm" yaml:"parm"`
}
