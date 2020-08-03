package infrastructure


import (
    "context"
   _ "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
    "sync"
    "log"
    _"github.com/LeonelSoriano/my-stock-framework/src/base/env"
    "strings"
)


type MongoDatabase struct {
    Client *mongo.Client
}

type MongoParams struct {
    User string
    Password string
    Host string
    Port string
    Database string
}

var MongoParamsLocal MongoParams;
//mongodb://root:rootpassword@localhost:27017/test
//mongodb://root:rootpassword@localhost:27017/admin

// Creador "estático"
var singleton *MongoDatabase
var once sync.Once
func GetMongoDatabase() *MongoDatabase {
     once.Do(func() {

        log.Printf(genenateUri(MongoParamsLocal))
        clientOptions := options.Client().
            ApplyURI(genenateUri(MongoParamsLocal))

        client, err := mongo.NewClient(clientOptions)
        if err != nil {
            log.Fatal(err)
        }

        ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
        err = client.Connect(ctx)
        if err != nil {
                log.Fatal(err)
        }

        err = client.Ping(context.TODO(), nil)

        if err != nil {
            log.Println(err)
        }

         singleton = &MongoDatabase{
             Client: client,
         }

        //defer client.Disconnect(ctx)
     })
     return singleton
}

func genenateUri (params MongoParams) string {

    //var str strings.Builder
    if strings.Compare(params.User, "") == 0 {
        panic("need add user")
    }

    if strings.Compare(params.Password, "") == 0 {
        panic("need add Password")
    }

    if strings.Compare(params.Host, "") == 0 {
        panic("need add Host")
    }

    if strings.Compare(params.Port, "") == 0 {
        panic("need add Port")
    }

    return "mongodb://" + params.User + ":" + params.Password + "@" +
        params.Host + ":" + params.Port + "/" + params.Database;
}
