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

/*import (
    "github.com/lnlwd/go-rest-ddd/domain"
    "log"
)

func main() {
    uf := domain.NewUserFile()
    users, err := uf.FindAllUsers()

    if err != nil {
        log.Println(err)
    }

    for _, user := range users {
        log.Println(user.Name)
    }
}*/