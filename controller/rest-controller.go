package controller

import (
	"encoding/json"
	"httpClientGo/model"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetResponse(ctx *gin.Context) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error. Check logs more info"})
		log.Printf("Error when request %s", err)
	}

	defer response.Body.Close()

	bodyResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error. Check logs more info"})
		log.Printf("Error read body %s", err)
	}

	log.Println("Body nya", string(bodyResponse))

	var postModel model.Post

	errMapEncodingJSON := json.Unmarshal(bodyResponse, &postModel)

	if errMapEncodingJSON != nil {
		log.Printf("Error when map json encoded %s", errMapEncodingJSON)
		ctx.JSON(500, gin.H{"error": "Internal server error. Check logs more info"})
	} else {
		serveResponse(ctx, postModel)
	}

}

func serveResponse(ctx *gin.Context, listPostData model.Post) {

	ctx.JSON(200, gin.H{
		"response": listPostData,
	})
}
