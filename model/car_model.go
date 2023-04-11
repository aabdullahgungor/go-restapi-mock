package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	Id    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Brand string             `bson:"brand,omitempty" json:"brand"`
	Series string             `bson:"series,omitempty" json:"series"`
	Year      time.Time     `bson:"year,omitempty" json:"year"`
	Fuel string             `bson:"fuel,omitempty" json:"fuel"`
	Gear string             `bson:"gear,omitempty" json:"gear"`
	Situation string             `bson:"situation,omitempty" json:"situation"`
	Km int	`bson:"km,omitempty" json:"km"`
	Color string             `bson:"color,omitempty" json:"color"`
	Price int	`bson:"price,omitempty" json:"price"`
}