package config

import (
	"encoding/json"
	"log"

	"github.com/apt-tool/apt-gateway/internal/config/core"
	"github.com/apt-tool/apt-gateway/internal/config/ftp"
	"github.com/apt-tool/apt-gateway/internal/config/http"
	"github.com/apt-tool/apt-gateway/internal/config/migration"
	"github.com/apt-tool/apt-gateway/internal/storage/redis"
	"github.com/apt-tool/apt-gateway/internal/storage/sql"
	"github.com/apt-tool/apt-gateway/internal/utils/jwt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/tidwall/pretty"
)

type Config struct {
	Core    core.Config      `koanf:"core"`
	HTTP    http.Config      `koanf:"http"`
	JWT     jwt.Config       `koanf:"jwt"`
	Redis   redis.Config     `koanf:"redis"`
	MySQL   sql.Config       `koanf:"mysql"`
	Migrate migration.Config `koanf:"migrate"`
	FTP     ftp.Config       `koanf:"ftp"`
}

func Load(path string) Config {
	var instance Config

	k := koanf.New(".")

	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	indent, err := json.MarshalIndent(instance, "", "\t")
	if err != nil {
		log.Fatalf("error marshaling config to json: %s", err)
	}

	indent = pretty.Color(indent, nil)
	tmpl := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(tmpl, string(indent))

	return instance
}