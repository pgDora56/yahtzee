package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "math/rand"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
    "gopkg.in/olahol/melody.v1"
)

var dices = [5]int {1,1,1,1,1}
var rollcnt = 0

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

    m.HandleMessage(func (s *melody.Session, msg []byte) {
        msgs := string(msg)
        log.Printf(msgs)
        if msgs == "Roll" {
            rollDice("01234")
            m.Broadcast(getDiceText())
        } else if strings.HasPrefix(msgs, "Reroll:") {
            rollDice(msgs[7:])
            m.Broadcast(getDiceText())
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

func rollDice(dicestr string) {
    if rollcnt == 0 {
        dicestr = "01234"
    }
    // rollcnt = (rollcnt + 1) % 3 // 実際の挙動
    rollcnt++ // デバック用
    rand.Seed(time.Now().UnixNano())
    for _, c := range dicestr {
        n := rand.Intn(6) + 1
        i := int(c - '0') // c : rune
        dices[i] = n
    }
}

func getDiceText() []byte {
    text := "R:"
    for _, i := range dices {
        text += " " + strconv.Itoa(i)
    }
    return []byte(text)
}

