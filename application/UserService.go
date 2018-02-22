package application

import (
    "github.com/gin-gonic/gin"
    "github.com/lnlwd/go-rest-ddd/domain"
)

func UsersList() func (c *gin.Context) {
    uf := domain.NewUserFile()

    users, err := uf.FindAllUsers()

    if err != nil {
        return StatusReport(500, err.Error())
    }

    x := gin.H{"data": users}

    return func(c *gin.Context) {
        c.JSON(200, x)
    }
}
