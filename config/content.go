package config

var config = new(Config)

type Config struct {
	App      App      `yaml:"app"`
	Jwt      Jwt      `yaml:"jwt"`
	Http     []Http   `yaml:"http"`
	Log      Log      `yaml:"log"`
	Redis    Redis    `yaml:"redis"`
	Database Database `yaml:"database"`
}

type App struct {
	Name     string `yaml:"name"`
	Debug    bool   `yaml:"debug"`
	Mode     string `yaml:"mode"`
	Timezone string `yaml:"timezone"`
}

type Jwt struct {
	Key             string `yaml:"key"`
	ExpireTime      int    `yaml:"expireTime"`
	MaxRefreshTime  int    `yaml:"maxRefreshTime"`
	DebugExpireTime int    `yaml:"debugExpireTime"`
}

type Http struct {
	Name              string `yaml:"name"`
	Port              string `yaml:"port"`
	TelemetryEndpoint string `yaml:"telemetryEndpoint"`
}

type Log struct {
	Level     string `yaml:"level"`
	MaxSize   int    `yaml:"maxSize"`
	MaxBackup int    `yaml:"maxBackup"`
	MaxAge    int    `yaml:"maxAge"`
	FilePath  string `yaml:"filePath"`
}

type Database struct {
	Connection string `yaml:"connection"`
	Mysql      struct {
		Host               string `yaml:"host"`
		Port               string `yaml:"port"`
		Username           string `yaml:"username"`
		Password           string `yaml:"password"`
		DBName             string `yaml:"dbname"`
		Charset            string `yaml:"charset"`
		MaxIdleConnections int    `yaml:"maxIdleConnections"`
		MaxOpenConnections int    `yaml:"maxOpenConnections"`
		MaxLifeSeconds     int    `yaml:"maxLifeSeconds"`
	} `yaml:"mysql"`
	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"postgres"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   int    `yaml:"dbname"`
}
