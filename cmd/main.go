package main

import (
	"fmt"

	"github.com/IcaroSilvaFK/object_orientation_on_go_lang/cmd/entities"
)

func main() {

	u := entities.NewUserEntity("Icaro", "Icaro@email", "password", 22)

	u.Save()

	fmt.Printf("%+v\n", u)

}
