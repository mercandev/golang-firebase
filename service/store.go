package service

import (
	"context"
	"encoding/json"
	"gofirestorre/domain"
	"io/ioutil"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Client() (*firestore.Client, context.Context) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./remote-mappers.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client, ctx
}

func GetData(w http.ResponseWriter, r *http.Request) {

	client, ctx := Client()

	iter := client.Collection("store").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		b, _ := json.Marshal(map[string]interface{}{
			"data": doc.Data(),
		})
		w.Write([]byte(b))
		//fmt.Println(doc.Data())
	}

}

func AddData(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var result domain.Store

	json.Unmarshal(reqBody, &result)

	client, ctx := Client()

	_, _, err := client.Collection("store").Add(ctx, result)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}
