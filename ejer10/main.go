package main

import (
	"fmt"
	"time"
	us "undemy/ejer10/user"
)

type Person struct {
	us.User
}

func main() {
	u := new(Person)
	u.AltaUser(1, "Ernesto Rodriguez", time.Now(), true)
	fmt.Println(u.User)

}
