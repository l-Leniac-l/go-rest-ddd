package domain

import (
    "github.com/lnlwd/go-rest-ddd/infra"
    "github.com/lnlwd/gremji"
    "encoding/json"
)

type UserEntity struct {
    Name string
}

func (userEntity UserEntity) GetAll() ([]UserEntity, error) {
    client := infra.ConnectToTinkerpop()
    res, err := client.ExecQuery(gremji.QueryArgs{"g.V().values('name')", nil, nil})

    if err != nil {
        return nil, err
    }

    data := res.Result.Data

    var datai map[string]interface{}

    if err = json.Unmarshal(data, &datai); err != nil {
        return nil, err
    }

    users := []UserEntity{}

    for _, name := range datai["@value"].([]interface{}) {
        user := UserEntity{Name: name.(string)}
        users = append(users, user)
    }

    return users, nil
}