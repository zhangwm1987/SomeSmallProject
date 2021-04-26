package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testConvert struct {
	ID   uuid.UUID `json:"_id" bson:"_id"`
	Name string    `json:"name" bson:"name"`
}

func main() {
	// Rest of the code will go here
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testcovertid").Collection("test")

	if len(os.Args) < 2 || (os.Args[1] != "UUID2BIN" && os.Args[1] != "BIN2UUID") {
		fmt.Println("Command should be format: \"mongouuid UUID2BIN/BIN2UUID xxxx\"")
		fmt.Println("BIN data only need the content, for example: BinData(0,\"9sfXXxQLRKWwwh0koojLjA==\")")
		fmt.Println("Pass in 9sfXXxQLRKWwwh0koojLjA==")
		return
	}
	if os.Args[1] == "UUID2BIN" {
		id, _ := uuid.FromString(os.Args[2])
		ash := testConvert{
			ID:   id,
			Name: "wenming",
		}
		_, err = collection.InsertOne(context.TODO(), ash)
		if err != nil {
			log.Fatal(err)
		}

		//out, err := exec.Command("mongo", "testcovertid", "<", "/home/zhangwm1987/workspace/test/mongodb/test.js").CombinedOutput()
		out, err := exec.Command("mongo", "testcovertid", "--eval", "db.test.find()").CombinedOutput()
		if err != nil {
			fmt.Println("wenming err")
			log.Fatal(err)
		}
		fmt.Println(string(out))

	} else {
		insertQuery := fmt.Sprintf("db.test.insert({ \"_id\" : BinData(0,\"%s\"), \"name\" : \"wenming\" })",
			os.Args[2])
		exec.Command("mongo", "testcovertid", "--eval", insertQuery).CombinedOutput()
		cur, _ := collection.Find(context.TODO(), bson.M{"name": "wenming"})
		for cur.Next(context.TODO()) {
			data := testConvert{}
			err := cur.Decode(&data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(data)
		}
	}
	collection.DeleteMany(context.TODO(), bson.M{"name": "wenming"})
}
