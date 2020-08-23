package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "math/rand"
    "strconv"

    "github.com/gin-gonic/gin"
    "gopkg.in/olahol/melody.v1"
)

func main() {
    log.Println("Websocket App Start!")
    router := gin.Default()
    router.Static("/assets", "./assets")
    m := melody.New()

    rg := router.Group("/sampleapp")
    rg.GET("/", func(ctx *gin.Context) {
        http.ServeFile(ctx.Writer, ctx.Request, "index.html")
    })

    rg.GET("/ws", func(ctx *gin.Context) {
        m.HandleRequest(ctx.Writer, ctx.Request)
    })

    m.HandleMessage(func(s *melody.Session, msg []byte) {
        msgs := string(msg)
        log.Printf(msgs)
        if msgs == "Roll" {
            rand.Seed(time.Now().UnixNano())
            msgs = "R:"
            for i := 0; i < 5; i++ {
                msgs += " "  + strconv.Itoa(rand.Intn(6) + 1)
            }
            m.Broadcast([]byte(msgs))
        } else {
            m.Broadcast(msg)
        }
    })

    m.HandleConnect(func(s *melody.Session) {
        log.Printf("websocket connection open, [session: %#v]\n", s)
    })

    m.HandleDisconnect(func(s *melody.Session) {
        log.Printf("websocket connection close. [session: %#v]\n", s)
    })

    router.Run(":51417")

    fmt.Println("WebSocket App End")
}
