package domain


type UserFiler interface{
    FindAllUsers() ([]UserEntity, error)
}

type userFile struct {

}

func (uf userFile) FindAllUsers() ([]UserEntity, error) {
    users, err := UserEntity{}.GetAll()
    if err != nil {
        return nil, err
    }
    return users, nil
}

func NewUserFile() UserFiler {
    return &userFile{}
}
