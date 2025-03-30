package migrations

import (
	"context"
	"log"

	"github.com/PavelBradnitski/Goods/auth_service/models"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Up() error {
	collection := mgm.Coll(&models.User{}).Collection

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Ошибка создания индекса: %v", err)
		return err
	}

	log.Println("Миграция выполнена: создан индекс для email")
	return nil
}

func Down() error {
	collection := mgm.Coll(&models.User{}).Collection
	_, err := collection.Indexes().DropOne(context.Background(), "email_1")
	if err != nil {
		log.Fatalf("Ошибка отката миграции: %v", err)
		return err
	}

	log.Println("Миграция отменена: индекс email удален")
	return nil
}
