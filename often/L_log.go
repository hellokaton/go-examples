package main

import (
	log1 "github.com/sirupsen/logrus"
	log2 "github.com/cihub/seelog"
	"os"
	"fmt"
)

func TestLogrus1()  {
	log1.WithFields(log1.Fields{
		"姓名": "王爵nice",
	}).Info("不小心中奖了")
}

func TestLogrus2()  {
	// 日志格式化为JSON而不是默认的ASCII
	log1.SetFormatter(&log1.JSONFormatter{})

	// 输出stdout而不是默认的stderr，也可以是一个文件
	log1.SetOutput(os.Stdout)

	// 只记录严重或以上警告
	log1.SetLevel(log1.WarnLevel)

	log1.WithFields(log1.Fields{
		"姓名": "王爵nice",
	}).Info("不小心没中奖╮(╯_╰)╭")

	log1.WithFields(log1.Fields{
		"姓名": "王爵nice",
	}).Error("不小心没中奖")
}

func TestSeeLog1()  {
	defer log2.Flush()
	log2.Info("Hello from Seelog!")
}

func TestSeeLog2()  {
	// xml格式的字符串，xml配置了如何输出日志
	testConfig := `
    <seelog type="sync">
        // 配置输出项，本例表示同时输出到控制台和日志文件中
    	<outputs formatid="main">
    		<console/>
    		<file path="log.log"/>
    	</outputs>
    	<formats>
    	    // 日志格式，outputs中的输出项可以通过id引用该格式
    		<format id="main" format="[%LEVEL] [%Time] [%FuncShort @ %File.%Line] %Msg%n"/>
    	</formats>
    </seelog>`

	// 根据配置信息生成logger（应该是配置信息的对象表示）
	logger, err := log2.LoggerFromConfigAsBytes([]byte(testConfig))
	if err != nil {
		fmt.Println(err)
	}
	// 应该是配置日志组件加载配置信息，然后输出函数比如Info在输出时，会加载该配置
	log2.ReplaceLogger(logger)
	log2.Info("Hello from Seelog!")
}
func main() {
	// logrus是用Go语言实现的一个日志系统，与标准库log完全兼容并且核心API很稳定,是Go语言目前最活跃的日志库
	TestLogrus1()
	TestLogrus2()

	// seelog是用Go语言实现的一个日志系统，它提供了一些简单的函数来实现复杂的日志分配、过滤和格式化
	TestSeeLog1()
	TestSeeLog2()
}