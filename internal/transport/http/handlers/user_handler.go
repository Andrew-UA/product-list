package handlers

import (
	"encoding/json"
	"github.com/Andrew-UA/product-list/app/services"
	"github.com/Andrew-UA/product-list/internal/transport/http/requests/user_requests"
	"github.com/Andrew-UA/product-list/internal/transport/http/responses"
	"github.com/Andrew-UA/product-list/internal/transport/http/responses/user_responses"
	"github.com/Andrew-UA/product-list/internal/validation"
	"net/http"
	"strconv"
)

type UserHandler struct {
	validator   validation.ValidatorInterface
	userService services.UsersServiceInterface
}

func NewUserHandler(validator validation.ValidatorInterface, userService services.UsersServiceInterface) *UserHandler {
	return &UserHandler{
		validator:   validator,
		userService: userService,
	}
}

func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, _ := h.userService.GetUsers(ctx)

	responses.NewResponseJson(http.StatusOK, user_responses.NewUserListResponse(users)).Write(w)
}

func (h *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdStr := r.PathValue("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		responses.NewResponseJson(http.StatusBadRequest, responses.NewErrorResponse(err)).Write(w)
		return
	}

	user, err := h.userService.GetUserById(ctx, uint64(userId))
	if err != nil {
		responses.NewResponseJson(http.StatusBadRequest, responses.NewErrorResponse(err)).Write(w)
		return
	}

	responses.NewResponseJson(http.StatusOK, user_responses.NewUserResponse(user)).Write(w)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request user_requests.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.NewResponseJson(http.StatusBadRequest, responses.NewErrorResponse(err)).Write(w)
		return
	}
	err = h.validator.ValidateStruct(&request)
	if err != nil {
		responses.NewResponseJson(http.StatusBadRequest, responses.NewErrorResponse(err)).Write(w)
		return
	}

	data := request.ToDTO()
	user, err := h.userService.CreateUser(r.Context(), data)
	if err != nil {
		responses.NewResponseJson(http.StatusInternalServerError, responses.NewErrorResponse(err)).Write(w)
		return
	}

	responses.NewResponseJson(http.StatusCreated, user_responses.NewUserResponse(user)).Write(w)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var request user_requests.UpdateUserRequest

	ctx := r.Context()
	userIdStr := r.PathValue("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		responses.NewResponseJson(http.StatusBadRequest, responses.NewErrorResponse(err)).Write(w)
	}

	user, err := h.userService.GetUserById(ctx, uint64(userId))
	if err != nil {
		responses.NewResponseJson(http.StatusBadRequest, responses.NewErrorResponse(err)).Write(w)
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.NewResponseJson(http.StatusBadRequest, responses.NewErrorResponse(err)).Write(w)
	}

	data := request.ToDTO()
	user, err = h.userService.UpdateUser(ctx, data, user)
	if err != nil {
		responses.NewResponseJson(http.StatusInternalServerError, responses.NewErrorResponse(err)).Write(w)
	}

	responses.NewResponseJson(http.StatusOK, user_responses.NewUserResponse(user)).Write(w)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userIdStr := r.PathValue("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		responses.NewResponseJson(http.StatusBadRequest, responses.NewErrorResponse(err)).Write(w)
		return
	}

	user, err := h.userService.GetUserById(ctx, uint64(userId))
	if err != nil {
		responses.NewResponseJson(http.StatusBadRequest, responses.NewErrorResponse(err)).Write(w)
		return
	}

	err = h.userService.DeleteUser(ctx, user)
	if err != nil {
		responses.NewResponseJson(http.StatusInternalServerError, responses.NewErrorResponse(err)).Write(w)
		return
	}

	responses.NewResponseJson(http.StatusOK, nil).Write(w)
}
