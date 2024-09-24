package handlers

import (
	"context"
	"project/internal/messagesService"
	"project/internal/web/messages"
)

type Handler struct {
	service *messagesService.MessageService
}

func (h *Handler) GetMessages(ctx context.Context, request messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	allMessages, err := h.service.FindAll()
	if err != nil {
		return nil, err
	}

	response := messages.GetMessages200JSONResponse{}

	for _, msg := range allMessages {
		message := messages.Message{
			Id:   &msg.ID,
			Text: &msg.Text,
		}
		response = append(response, message)
	}

	return response, nil
}

func (h *Handler) PostMessages(ctx context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body

	messageToCreate := messagesService.Message{Text: messageRequest.Text}
	createdMessage, err := h.service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}

	response := messages.PostMessages201JSONResponse{
		Id:   &createdMessage.ID,
		Text: &createdMessage.Text,
	}

	return response, nil
}

func (h *Handler) PatchMessagesId(ctx context.Context, request messages.PatchMessagesIdRequestObject) (messages.PatchMessagesIdResponseObject, error) {
	messageToUpdate := messagesService.Message{Text: request.Body.Text}

	updatedMessage, err := h.service.UpdateMessageByID(request.Id, messageToUpdate)
	if err != nil {
		if err.Error() == "not found" {
			return messages.PatchMessagesId404Response{}, nil
		}
		return nil, err
	}

	response := messages.PatchMessagesId200JSONResponse{
		Id:   &updatedMessage.ID,
		Text: &updatedMessage.Text,
	}

	return response, nil
}

func (h *Handler) DeleteMessagesId(ctx context.Context, request messages.DeleteMessagesIdRequestObject) (messages.DeleteMessagesIdResponseObject, error) {
	err := h.service.DeleteMessageByID(request.Id)
	if err != nil {
		if err.Error() == "not found" {
			return messages.DeleteMessagesId404Response{}, nil
		}
		return nil, err
	}

	return messages.DeleteMessagesId204Response{}, nil
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{service: service}
}
