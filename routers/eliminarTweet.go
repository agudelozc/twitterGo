package routers

import (
	"context"

	"github.com/agudelozc/twitterGo/bd"
	"github.com/agudelozc/twitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

func EliminarTweet(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El párametro ID es obligatorio"
		return r
	}
	err := bd.BorroTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrio un error al intentar borrar el tweet" + err.Error()
		return r
	}

	r.Message = "Eliminar Tweet Ok!"
	r.Status = 200
	return r

}
