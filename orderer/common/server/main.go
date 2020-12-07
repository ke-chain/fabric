package server

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/ke-chain/fabric/orderer/common/localconfig"
	"github.com/ke-chain/fabric/orderer/common/server/metadata"
	"gopkg.in/alecthomas/kingpin.v2"
)

var logger = flogging.MustGetLogger("orderer.common.server")

//command line flags
var (
	app = kingpin.New("orderer", "ke-chain Fabric orderer node")

	_       = app.Command("start", "Start the orderer node").Default() // preserved for cli compatibility
	version = app.Command("version", "Show version information")
)

// Main is the entry point of orderer process
func Main() {
	fullCmd := kingpin.MustParse(app.Parse(os.Args[1:])) // 解析用户命令行

	// "version" command
	if fullCmd == version.FullCommand() {
		fmt.Println(metadata.GetVersionInfo())
		return
	}

	conf, err := localconfig.Load() //  记载orderer配置文件
	if err != nil {
		logger.Error("failed to parse config: ", nil)
		os.Exit(1)
	}
	initializeLogging()

	prettyPrintStruct(conf)
}

func initializeLogging() {
	loggingSpec := os.Getenv("FABRIC_LOGGING_SPEC")
	loggingFormat := os.Getenv("FABRIC_LOGGING_FORMAT")
	flogging.Init(flogging.Config{
		Format:  loggingFormat,
		Writer:  os.Stderr,
		LogSpec: loggingSpec,
	})
}

func prettyPrintStruct(i interface{}) {
	params := localconfig.Flatten(i)
	var buffer bytes.Buffer
	for i := range params {
		buffer.WriteString("\n\t")
		buffer.WriteString(params[i])
	}
	logger.Infof("Orderer config values:%s\n", buffer.String())
}
