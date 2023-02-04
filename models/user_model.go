package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `JSON:"id,omitempty"`
	Name     string             `JSON:"name, omitempty"`
	Email    string             `JSON:"email, omitempty" validate:"required"`
	Password string             `JSON:"password, omitempty" validate:"required"`
}
