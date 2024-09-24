package handlers

import (
	"context"
	"project/internal/userService"
	"project/internal/web/users"
)

// UserHandler handles requests related to users.
type UserHandler struct {
	service *userService.UserService
}

// GetUsers fetches all users.
func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:        &usr.ID,
			Email:     &usr.Email,
			Password:  &usr.Password,
			CreatedAt: &usr.CreatedAt,
			UpdatedAt: &usr.UpdatedAt,
		}
		response = append(response, user)
	}

	return response, nil
}

// PostUsers creates a new user.
func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	createdUser, err := h.service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:        &createdUser.ID,
		Email:     &createdUser.Email,
		Password:  &createdUser.Password,
		CreatedAt: &createdUser.CreatedAt,
		UpdatedAt: &createdUser.UpdatedAt,
	}

	return response, nil
}

// PatchUsersId updates an existing user by ID.
func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userToUpdate := userService.User{
		Email:    request.Body.Email,
		Password: request.Body.Password,
	}

	updatedUser, err := h.service.UpdateUserByID(request.Id, userToUpdate)
	if err != nil {
		if err.Error() == "not found" {
			return users.PatchUsersId404Response{}, nil
		}
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:        &updatedUser.ID,
		Email:     &updatedUser.Email,
		Password:  &updatedUser.Password,
		CreatedAt: &updatedUser.CreatedAt,
		UpdatedAt: &updatedUser.UpdatedAt,
	}

	return response, nil
}

// DeleteUsersId deletes a user by ID.
func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.service.DeleteUserByID(request.Id)
	if err != nil {
		if err.Error() == "not found" {
			return users.DeleteUsersId404Response{}, nil
		}
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{service: service}
}
