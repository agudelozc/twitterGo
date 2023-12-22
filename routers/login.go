package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/agudelozc/twitterGo/bd"
	"github.com/agudelozc/twitterGo/jwt"
	"github.com/agudelozc/twitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.RespApi {
	var r models.RespApi
	var t models.Usuario
	r.Status = 400
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "Email requerido"
		return r
	}
	userData, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		r.Message = "Usuario o contraseña incorrectos"
		return r
	}

	jwtKey, err := jwt.GeneroJWT(ctx, userData)
	if err != nil {
		r.Message = "Ocurrio un erro al intentar generar el token" + err.Error()
		return r
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	token, err2 := json.Marshal(resp)
	if err2 != nil {
		r.Message = "Ocurrio un erro al intentar formatear el token a json" + err2.Error()
		return r
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}
	cookieString := cookie.String()

	res:= &events.APIGatewayProxyResponse{
		Headers: map[string]string{
            "Content-Type": "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie": cookieString,
        },
        Body: string(token),
        StatusCode: 200,
	}
	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res
	return r

}
