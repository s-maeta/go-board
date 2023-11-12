package config

import (
	"fmt"
	"os"
	"regexp"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// config設定の構造体
type Config struct {
	Env     string
	Tz      string
	ApiUrl  string
	Db      Db      `yml:db`
	Mail    Mail    `yml:mail`
	Migrate Migrate `yml:migrate`
	Auth    `yml:auth`
}

type Db struct {
	Type      string `yml:type`
	Host      string `yml:host`
	Port      int    `yml:port`
	Charset   string `yml:charset`
	ParseTime bool   `yml:parseTime`
	Loc       string `yml:loc`
	Database  string
	User      string
	Password  string
}

type Mail struct {
	Host string `yml:host`
	Port int    `yml:port`
}

type Migrate struct {
	FilePath string `yml:filePath`
}

type Auth struct {
	SecretKey     string `yml:secretKey`
	TokenLifetime int    `yml:tokenLifetime`
}

var cfg *Config

func Init() {
	const projectDirName = "backend"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	// err := godotenv.Load(string(rootPath) + "/.env")
	// if  err != nil {
	// 	panic(err)
	// }

	// viperの初期設定
	viper.SetConfigName("config_" + fmt.Sprintf("%s", os.Getenv("ENV")))
	viper.SetConfigType("yml")
	viper.AddConfigPath(string(rootPath) + "/src/config/")

	// configファイル更新時に再読み込み
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		viper.Unmarshal(&cfg)
	})

	// configファイルの読み込み
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 読み込んだデータを変数cfgに設定
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	// 環境変数の値を変数cfgに設定
	cfg.Env = os.Getenv("ENV")
	cfg.Tz = os.Getenv("TZ")
	cfg.Db.Database = os.Getenv("MYSQL_DATABASE")
	cfg.Db.User = os.Getenv("MYSQL_USER")
	cfg.Db.Password = os.Getenv("MYSQL_ROOT_PASSWORD")
	cfg.ApiUrl = os.Getenv("API_URL")
}

// Configを取得する
func GetConfig() *Config {
	return cfg
}
