package models

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string    `json:"username" bson:"username"`
	Password         string    `json:"password" bson:"password"`
	Email            string    `json:"email" bson:"email"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}

// Коллекция для пользователей
func (u *User) CollectionName() string {
	return "users"
}

// Create индексирует коллекцию пользователей
func (u *User) Create() error {
	collection := mgm.Coll(u)
	indexView := collection.Indexes()

	model := mongo.IndexModel{
		Keys: bson.D{{Key: "username", Value: 1}},
		Options: &options.IndexOptions{
			Name:   &[]string{"username_index"}[0], // Необходимо передать указатель на строку
			Unique: &[]bool{true}[0],               // Необходимо передать указатель на bool
		},
	}

	_, err := indexView.CreateOne(context.Background(), model)
	return err
}

// Перед сохранением хешируем пароль
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

// Перед обновлением обновляем время изменения
func (u *User) BeforeUpdate(ctx context.Context) error {
	u.UpdatedAt = time.Now()
	return nil
}

// VerifyPassword проверяет пароль
func (u *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// FindByID ищет пользователя по ID
func FindByID(id string) (*User, error) {
	user := &User{}
	err := mgm.Coll(user).FindByID(id, user)
	return user, err
}

// FindByUsername ищет пользователя по username
func FindByUsername(username string) (*User, error) {
	user := &User{}
	err := mgm.Coll(user).FindOne(context.Background(), bson.M{"username": username}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, nil // Пользователь не найден, возвращаем nil
	}

	if err != nil {
		return nil, err // Ошибка при запросе
	}
	return user, nil
}
