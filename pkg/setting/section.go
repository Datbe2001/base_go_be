package setting

type Config struct {
	Mysql  MySQLSetting  `map_structure:"mysql"`
	Logger LoggerSetting `map_structure:"logger"`
}

type MySQLSetting struct {
	Host            string `map_structure:"host"`
	Port            int    `map_structure:"port"`
	UserName        string `map_structure:"username"`
	Password        string `map_structure:"password"`
	DBName          string `map_structure:"dbname"`
	MaxIdleConn     int    `map_structure:"maxIdleConn"`
	MaxOpenConn     int    `map_structure:"maxOpenConn"`
	ConnMaxLifeTime int    `map_structure:"connMaxLifeTime"`
}

type LoggerSetting struct {
	Log_level     string `map_structure:"log_level"`
	File_log_name string `map_structure:"file_log_name"`
	Max_backups   int    `map_structure:"max_backups"`
	Max_size      int    `map_structure:"max_size"`
	Max_age       int    `map_structure:"max_age"`
	Compress      bool   `map_structure:"compress"`
}
