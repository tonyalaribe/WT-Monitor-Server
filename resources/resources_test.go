package resources

import (
	"flag"
	"os"
	"testing"

	"github.com/tonyalaribe/monitor-server/config"
	"github.com/tonyalaribe/monitor-server/logger"
)

func TestMain(m *testing.M) {
	configFile := flag.String("config", "../config.toml", "path to config file")
	flag.Parse()
	config.Init(*configFile)
	logger.Init()

	retCode := m.Run()

	os.Exit(retCode)
}
