package server

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/ke-chain/fabric/orderer/common/localconfig"
	"github.com/ke-chain/fabric/orderer/common/server/metadata"
	"go.uber.org/zap"
	"gopkg.in/alecthomas/kingpin.v2"
)

var logger *zap.SugaredLogger

func init() {
	pro, _ := zap.NewProduction()
	defer pro.Sync() // flushes buffer, if any
	logger = pro.Sugar()
}

//command line flags
var (
	app = kingpin.New("orderer", "ke-chain Fabric orderer node")

	_       = app.Command("start", "Start the orderer node").Default() // preserved for cli compatibility
	version = app.Command("version", "Show version information")
)

// Main is the entry point of orderer process
func Main() {
	fullCmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	// "version" command
	if fullCmd == version.FullCommand() {
		fmt.Println(metadata.GetVersionInfo())
		return
	}

	conf, err := localconfig.Load()
	if err != nil {
		logger.Error("failed to parse config: ", err)
		os.Exit(1)
	}
	spew.Dump(conf)
}
