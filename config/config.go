package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

type Config struct {
	IsProduction bool   `toml:"is_production"`
	Domain       string `toml:"domain"`
	Port         string `toml:"port"`
	StaticDir    string `toml:"static_dir"`

	LogFile string `toml:"log_file"`

	PasswordEncryptionCost int `toml:"password_encryption_cost"`

	Cors struct {
		AllowedOrigins []string `toml:"allowed_origins"`
	}

	AWS struct {
		Region         string
		ImageBucket    string `toml:"image_bucket"`
		ResumeBucket   string `toml:"resume_bucket"`
		ProfileRoot    string `toml:"profile_root"`
		UserDefault    string `toml:"user_default"`
		CompanyDefault string `toml:"company_default"`
	}

	DB struct {
		File     string `toml:"file"`
		TestFile string `toml:"test_file"`
	}

	Encryption struct {
		Private []byte
		Public  []byte
	}
}

var (
	config     Config
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func Get() *Config {
	return &config
}

func Init(file string) {
	_, err := os.Stat(file)
	if err != nil {
		log.Fatal("Config file is missing: ", file)
	}
	if _, err := toml.DecodeFile(file, &config); err != nil {
		log.Fatal(err)
	}

	log.Println(basepath)
	config.Encryption.Public, err = ioutil.ReadFile("./config/encryption_keys/public.pem")
	if err != nil {
		log.Println("Error reading public key")
		log.Println(err)
		return
	}

	config.Encryption.Private, err = ioutil.ReadFile("./config/encryption_keys/private.pem")
	if err != nil {
		log.Println("Error reading private key")
		log.Println(err)
		return
	}
}
