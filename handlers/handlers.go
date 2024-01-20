package handlers

import (
	"context"
	"fmt"

	"github.com/agudelozc/twitterGo/jwt"
	"github.com/agudelozc/twitterGo/models"
	"github.com/agudelozc/twitterGo/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Voy a procesar" + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespApi
	r.Status = 400

	isOk, statusCode, msg, claim := validoAuhorization(ctx, request)

	if !isOk {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "verperfil":
			//return routers.VerPerfil(request)
		}
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)
		case "login":
			return routers.Login(ctx)
		case "tweet":
			return routers.GraboTweet(ctx, claim)
		}
		
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		//
		case "modificarPerfil":
			return routers.ModificarPerfil(ctx, claim )
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		case "/":
		}

	}
	r.Message = "Method Invalid"
	return r
}

func validoAuhorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}
	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	claim, todoOK, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtsign")).(string))

	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token" + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token" + msg)
			return false, 401, msg, models.Claim{}
		}

	}
	fmt.Println("Token OK")
	return true, 200, msg, *claim
}
