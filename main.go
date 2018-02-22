package main

import (
    "github.com/gin-gonic/gin"
    "github.com/lnlwd/go-rest-ddd/application"
)

func main() {
    router := gin.Default()
    application.HandleRequests(router)
    router.Run()
}