package config

import (
	"context"
	"os"

	"github.com/MariliaNeves/api-genealogy/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL          = "MONGODB_URL"
	MONGODB_GENEALOGY_DB = "MONGODB_GENEALOGY_DB"
	MONGODB_USERNAME     = "MONGODB_USERNAME"
	MONGODB_PASSWORD     = "MONGODB_PASSWORD"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {

	mongodbURI := os.Getenv(MONGODB_URL)
	mongodbDatabase := os.Getenv(MONGODB_GENEALOGY_DB)
	mongodbUsername := os.Getenv(MONGODB_USERNAME)
	mongodbPassword := os.Getenv(MONGODB_PASSWORD)

	clientOptions := options.Client().ApplyURI(mongodbURI)
	if mongodbUsername != "" && mongodbPassword != "" {
		clientOptions.Auth = &options.Credential{
			Username: mongodbUsername,
			Password: mongodbPassword,
		}
	}

	// Conectar ao MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error("Error trying to connect to MongoDB", err)
	}

	// Verificar conexão
	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Error trying to PING in MongoDB", err)
	}

	// Retornar o banco de dados
	return client.Database(mongodbDatabase), nil
}
