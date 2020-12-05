package server

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/ke-chain/fabric/orderer/common/localconfig"
	"github.com/ke-chain/fabric/orderer/common/server/metadata"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/alecthomas/kingpin.v2"
)

var logger *zap.SugaredLogger

func init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./foo.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger = zap.New(core).Sugar()
}

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
	logger.Info(spew.Sdump(conf)) // log 记录所有配置信息
	logger.Sync()                 // 输出 log 缓存
}
