package infra

import (
    "log"
    "github.com/lnlwd/gremji"
)

func ConnectToTinkerpop() *gremji.Client {
    client, err := gremji.NewClient("ws://localhost:8182/gremlin")
    if err != nil {
        log.Println(err)
        return nil
    }

    return client
}