package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevueltoTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"_userid" json:"_userid,omitempty"`
	Mensaje string             `bson:"_mensaje" json:"_mensaje,omitempty"`
	Fecha   time.Time          `bson:"_fecha" json:"_fecha,omitempty"`
}
