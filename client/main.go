package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hidori/go-test-openapi/contactsapi"
	"github.com/pkg/errors"
)

func main() {
	client, err := contactsapi.NewClientWithResponses("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	getContactList(ctx, client)
	addContact(ctx, client, "firstName1", "familyName1", "1111111111")
	addContact(ctx, client, "firstName2", "familyName2", "2222222222")
	getContactList(ctx, client)
}

func getContactList(ctx context.Context, client *contactsapi.ClientWithResponses) {
	response, err := client.GetContactListWithResponse(ctx, &contactsapi.GetContactListParams{})
	if err != nil {
		log.Fatal(err)
	}

	s, err := toJSONString(response.JSON200)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("---- getContactList\n%s\n\n", s)
}

func addContact(ctx context.Context, client *contactsapi.ClientWithResponses, firstName string, familyName string, phoneNumber string) {
	response, err := client.AddContactWithResponse(ctx, contactsapi.ContactValues{
		FirstName:   firstName,
		FamilyName:  familyName,
		PhoneNumber: phoneNumber,
	})
	if err != nil {
		log.Fatal(err)
	}

	s, err := toJSONString(response.JSON201)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("---- addContact\n%s\n\n", s)
}

func toJSONString(object any) (string, error) {
	bytes, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(bytes), nil
}
