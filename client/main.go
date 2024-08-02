package main

import (
	"context"
	"log"

	"github.com/hidori/go-test-openapi/contactsapi"
)

func main() {
	client, err := contactsapi.NewClient("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	response, err := client.GetContactList(ctx, &contactsapi.GetContactListParams{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
