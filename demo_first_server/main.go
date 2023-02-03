package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// 普通gin
	//r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "Hello, Geektutu")
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080

	// websocket
	server := NewRouter()
	server.Run()
}

func RunSocekt(c *gin.Context) {
	// 可以接受socket参数
	user := c.Query("user")
	if user == "" {
		return
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
BYT:
	_, message, err := ws.ReadMessage()
	if nil != err {
		fmt.Println(err.Error())
	}
	ws.WriteMessage(websocket.TextMessage, message)
	goto BYT
}

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(Cors())

	socket := RunSocekt

	group := server.Group("")
	{
		//group.GET("/user", v1.GetUserList)
		group.GET("/socket", socket)
	}
	return server
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		c.Next()
	}
}
