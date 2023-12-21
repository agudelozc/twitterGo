package routers

import (
    "context"
    "fmt"
	"encoding/json"
	"github.com/agudelozc/twitterGo/bd"
	"github.com/agudelozc/twitterGo/models"
)

func Registro(ctx context.Context) models.RespApi {
	var r models.RespApi
	var t models.Usuario
	r.Status = 400
	fmt.Println("Entre al registro")
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err!= nil {
        r.Message = err.Error()
		fmt.Println(r.Message)
        return r
    }
	if len(t.Email) == 0 {
		r.Message = "Email requerido"
		fmt.Println(r.Message)
        return r
	}
	if len(t.Password) < 6 {
		r.Message = "La contraseña debe tener al menos 6 caracteres"
		fmt.Println(r.Message)
        return r
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
        r.Message = "Ya existe un usuario con ese email"
        fmt.Println(r.Message)
        return r
    }
	_, status, err := bd.InsertoRegistro(t)
	if err!= nil {
        r.Message = "Ocurrio un error al intentar realizar el registro de usuario"+err.Error()
        fmt.Println(r.Message)
        return r
    }
	if !status {
		r.Message = "No se ha logrado insertar el registro del usuario"
		fmt.Println(r.Message)
        return r

	}

	r.Status = 200
	r.Message = "Registro exitoso"
	fmt.Println(r.Message)
	return r

}