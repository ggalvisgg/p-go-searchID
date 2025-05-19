package repositories

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "example.com/go-mongo-app/models"
    "log"
    "os"  

)


type BookRepository struct {
    collection *mongo.Collection
}

func NewBookRepository() *BookRepository {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI not set in environment")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("library").Collection("books")
    return &BookRepository{collection}
}


func (r *BookRepository) GetBookByID(id string) (*models.Book, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var book models.Book
    filter := bson.M{"_id": objectID}
    err = r.collection.FindOne(context.TODO(), filter).Decode(&book)
    if err != nil {
        return nil, err
    }

    return &book, nil
}
