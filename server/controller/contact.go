package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/hidori/go-test-openapi/contactsapi"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/maps"
)

type ContactController struct {
	contacts map[string]contactsapi.Contact
}

var controller = &ContactController{
	contacts: map[string]contactsapi.Contact{},
}

func GetContactController() contactsapi.ServerInterface {
	return controller
}

func (s *ContactController) GetContactList(ctx echo.Context, params contactsapi.GetContactListParams) error {
	response := &contactsapi.ContactList{
		Items: maps.Values(controller.contacts),
	}

	return ctx.JSON(http.StatusOK, &response)
}

func (s *ContactController) AddContact(ctx echo.Context) error {
	var requestBody contactsapi.AddContactJSONRequestBody

	err := ctx.Bind(&requestBody)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "BadRequest")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "InternalServerError")
	}

	_id := id.String()

	response := contactsapi.Contact{
		Id:          _id,
		FamilyName:  requestBody.FamilyName,
		FirstName:   requestBody.FirstName,
		PhoneNumber: requestBody.PhoneNumber,
	}

	controller.contacts[_id] = response

	return ctx.JSON(http.StatusCreated, response)
}

func (s *ContactController) DeleteContactById(ctx echo.Context, id string) error {
	delete(controller.contacts, id)

	return ctx.String(http.StatusOK, "OK")
}

func (s *ContactController) GetContactById(ctx echo.Context, id string) error {
	response, ok := controller.contacts[id]
	if !ok {
		return ctx.String(http.StatusNotFound, "NotFound")
	}

	ctx.JSON(http.StatusOK, response)

	return nil
}
