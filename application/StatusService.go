package application

import "github.com/gin-gonic/gin"

func StatusReport(code int, status string) func(c *gin.Context) {
    return func(c *gin.Context) {
        c.JSON(code, gin.H{
            "status": status,
        })
    }
}