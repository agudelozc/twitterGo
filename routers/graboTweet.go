package routers

import (
	"context"
	"encoding/json"
	"time"
	"github.com/agudelozc/twitterGo/bd"
	"github.com/agudelozc/twitterGo/models"
)

func GraboTweet(ctx context.Context, claim models.Claim) models.RespApi{
	var mensaje models.Tweet

	var r models.RespApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensaje)
	if err != nil{
		r.Message="Ocurrio un error al intentar decodificar el body"+err.Error()
		return r
	}
	registro := models.GraboTweet{
		UserID: IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha: time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil{
		r.Message="Ocurrio un error al insertar el registro"+err.Error()
		return r
	}
	if !status{
		r.Message="No se ha logrado insertar el tweet"
		return r
	}
	r.Status=200
	r.Message="Tweet creado correctamente"
	return r
}