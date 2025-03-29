package models

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	mgm.DefaultModel `swaggerignore:"true" bson:",inline"`
	Username         string    `json:"username" bson:"username"`
	Password         string    `json:"password" bson:"password"`
	Email            string    `json:"email" bson:"email"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}

func (u *User) BeforeCreate(ctx context.Context) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(ctx context.Context) error {
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func FindByID(id string) (*User, error) {
	user := &User{}
	err := mgm.Coll(user).FindByID(id, user)
	return user, err
}

func FindByUsername(username string) (*User, error) {
	user := &User{}
	err := mgm.Coll(user).FindOne(context.Background(), bson.M{"username": username}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return user, nil
}
