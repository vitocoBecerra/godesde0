package usuarios

import (
	"fmt"
	"godesde0/modelos"
	"time"
)

func AltaUsuario() {

	u := new(modelos.User)
	u.AddUser(1, "Vitoco", time.Now(), true)
	fmt.Println(u)

}
