package models

import (
	"flag"
	"os"
	"testing"

	"github.com/tonyalaribe/monitor-server/config"
)

func TestMain(m *testing.M) {
	configFile := flag.String("config", "../config.toml", "path to config file")
	flag.Parse()
	config.Init(*configFile)

	retCode := m.Run()

	os.Exit(retCode)
}
