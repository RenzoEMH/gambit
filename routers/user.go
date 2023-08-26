package routers

import (
	"encoding/json"

	"github.com/renzoemh/gambit/bd"
	"github.com/renzoemh/gambit/models"
)

func UpdateUser(body string, User string) (int, string) {
	var t models.User
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.UserFirstName) == 0 && len(t.UserLastName) == 0 {
		return 400, "Debe especificar el Nombre (Firstname) o (LastName) del usuario"
	}

	_, encontrado := bd.UserExists(User)
	if !encontrado {
		return 400, "No existe un usuario con ese UUID '" + User + "'"
	}

	err = bd.UpdateUser(t, User)
	if err != nil {
		return 400, "Ocurrioo un error al intentar realizar la actualizacion del usuario " + User + " > " + err.Error()
	}

	return 200, "UpdateUser Ok"
}
