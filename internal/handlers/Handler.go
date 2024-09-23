package handlers

import (
	"context"
	"project/internal/messagesService"
	"project/internal/web/messages"
)

type Handler struct {
	service *messagesService.MessageService
}

func (h *Handler) GetGet(ctx context.Context, request messages.GetGetRequestObject) (messages.GetGetResponseObject, error) {
	allMessages, err := h.service.FindAll()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := messages.GetGet200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, msg := range allMessages {
		message := messages.Message{
			Id:   &msg.ID,
			Text: &msg.Text,
		}
		response = append(response, message)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostPost(ctx context.Context, request messages.PostPostRequestObject) (messages.PostPostResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := messagesService.Message{Text: messageRequest.Text}
	createdMessage, err := h.service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PostPost201JSONResponse{
		Id:   &createdMessage.ID,
		Text: &createdMessage.Text,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) PatchPatchId(ctx context.Context, request messages.PatchPatchIdRequestObject) (messages.PatchPatchIdResponseObject, error) {
	messageToUpdate := messagesService.Message{Text: request.Body.Text}
	updatedMessage, err := h.service.UpdateMessageByID(request.Id, messageToUpdate)

	if err != nil {
		if err.Error() == "not found" {
			return messages.PatchPatchId404Response{}, nil
		}
		return nil, err
	}

	response := messages.PatchPatchId200JSONResponse{
		Id:   &updatedMessage.ID,
		Text: &updatedMessage.Text,
	}
	return response, nil
}

func (h *Handler) DeleteDeleteId(ctx context.Context, request messages.DeleteDeleteIdRequestObject) (messages.DeleteDeleteIdResponseObject, error) {
	err := h.service.DeleteMessageByID(request.Id)

	if err != nil {
		if err.Error() == "not found" {
			return messages.DeleteDeleteId404Response{}, nil
		}
		return nil, err
	}
	return messages.DeleteDeleteId204Response{}, nil
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{service: service}
}
