package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type photo struct {
    ID string `json:"id"`
    UserID string `json:"user_id"`
    PostDate string `json:"post_date"`
    Tag []string  `json:"tag"`
    URL string `json:"url"`
}

var photos = []photo{
    {ID: "00000001", UserID: "0001", PostDate: "2022.04.20", Tag: []string{"クレープ", "いちご"}, URL: "https://cdn.discordapp.com/attachments/755257997789233195/990916146825920522/DSC_00012_23.jpeg"},
    {ID: "00000002", UserID: "0001", PostDate: "2022.05.11", Tag: []string{"トースト", "あずき"}, URL: "https://cdn.discordapp.com/attachments/755257997789233195/990916205638484038/DSC_00012_15.jpeg"},
    {ID: "00000003", UserID: "0001", PostDate: "2022.06.20", Tag: []string{"生ハム", "ワイン", "白"}, URL: "https://cdn.discordapp.com/attachments/755257997789233195/990916206804488222/DSC_0200.jpg"},
    {ID: "00000004", UserID: "0001", PostDate: "2022.06.21", Tag: []string{"海鮮丼", "サーモン", "いくら"}, URL: "https://cdn.discordapp.com/attachments/755257997789233195/990916227989897226/DSC_0064.jpg"},
}

func main() {
    router := gin.Default()

    config := cors.DefaultConfig()
    // config.AllowOrigins = []string{"http://sample.com"}
    config.AllowOrigins = []string{"*"}
    router.Use(cors.New(config))

    router.GET("/photos", getPhotos)
    router.GET("/", getPhotos)

    router.POST("post", postFunc)

    router.Run("localhost:8080")
}

func getPhotos(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, photos)
}

type InputCompany struct {
    Name string
}

func postFunc(c *gin.Context) {
    var hoge InputCompany
    c.BindJSON(&hoge)
    fmt.Println(hoge.Name)

    c.JSON(200, gin.H{
        "message": hoge.Name + "さん、こんにちは",
    })
}
