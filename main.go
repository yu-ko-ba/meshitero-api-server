package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

type photo struct {
    ID string `json:"id"`
    UserID string `json:"user_id"`
    PostDate string `json:"post_date"`
    Tag string `json:"tag"`
    URL string `json:"url"`
}

var photos = []photo{
    {ID: "00000001", UserID: "0001", PostDate: "2022.04.20", Tag: "#りんご,#山形りんご", URL: "https://columbia.jp/idolmaster/img/COCC-18008.jpg"},
    {ID: "00000002", UserID: "0001", PostDate: "2022.04.20", Tag: "#すだち,#徳島すだち", URL: "https://columbia.jp/idolmaster/img/COCC-18009.jpg"},
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
