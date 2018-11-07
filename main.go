package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"litilegame/config"
	"litilegame/dbconnect"
	"litilegame/logger"
	"net/http"
	"os"
)



var(
	configPath string
	help	bool

)
func init() {
	flag.BoolVar(&help,"h",false,"show help")
	flag.StringVar(&configPath,"c","","define config.ini file path when you use command line way to " +
		"start server")
	flag.BoolVar(&logger.Debug,"d",false,"when true is running in debug model")
	flag.Parse()
	flag.Usage=usage

	if configPath=="" {
		panic("配置文件参数不能为空，请指定-c参数，更多参数查看请使用-h")
	}
	fmt.Printf("系统初始化中,正在从,%s,配置文件中读取配置\n",configPath)
	config.LoadConfigFromConfigINI(configPath)
	fmt.Println("配置文件加载成功")
	fmt.Println("正在初始化日志器")
	logger.InitLog()
	fmt.Println("日志系统初始化完毕")
	fmt.Println("正在初始化mysql连接池")
	dbconnect.InitMysql()
	fmt.Println("mysql连接池初始化完毕")
	fmt.Println("正在初始化redis连接池")
	dbconnect.InitRedis()
	fmt.Println("redis连接池初始化完毕")
	fmt.Println("系统初始化完毕")


}



var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

//webSocket 请求 ping 返回 pong
func ping(c *gin.Context) {
	// 升级 get 请求为 webSocket 协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		// 读取 ws 中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		// 写入 ws 数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		//c.Set("request", "clinet_request")
		c.Next()
		fmt.Println("after middleware")
	}
}

func main() {
	bindAddress := "localhost:2303"
	r := gin.Default()
	r.GET("/ping", ping)
	r.Use(MiddleWare())
	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("hello")
		c.JSON(200,"hello middleware")
	})
	r.POST("/register", func(c *gin.Context) {

	})
	r.Run(bindAddress)
}

func usage()  {
	fmt.Fprintf(os.Stderr, `tianqisever version: 1.0
Usage: tianqiserver  [-d Debug] [-c configpath] [-fws freshweatehrspeed] [-cws coolweatehrspeed] [-fas freshalarmspeed] [-cas coolalarmspeed]

Options:
`)
	flag.PrintDefaults()
}