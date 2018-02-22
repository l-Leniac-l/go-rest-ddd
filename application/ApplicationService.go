package application

import "github.com/gin-gonic/gin"

func HandleRequests(router gin.IRouter) {
    router.GET("/", StatusReport(200,"Running"))
    router.GET("/users", UsersList())
}