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

	user_responses.NewUserListResponse(users).SetStatusCode(http.StatusOK).Write(w)
}

func (h *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdStr := r.URL.Query().Get("id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	user, _ := h.userService.GetUserById(ctx, userId)

	user_responses.NewUserResponse(user).SetStatusCode(http.StatusOK).Write(w)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request user_requests.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = h.validator.ValidateStruct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	data := request.ToDTO()
	user, _ := h.userService.CreateUser(r.Context(), data)

	user_responses.NewUserResponse(user).SetStatusCode(http.StatusCreated).Write(w)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var request user_requests.UpdateUserRequest

	ctx := r.Context()
	userIdStr := r.URL.Query().Get("id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	user, _ := h.userService.GetUserById(ctx, userId)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	data := request.ToDTO()
	user, _ = h.userService.UpdateUser(ctx, data, user)

	user_responses.NewUserResponse(user).SetStatusCode(http.StatusOK).Write(w)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userIdStr := r.URL.Query().Get("id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	user, _ := h.userService.GetUserById(ctx, userId)

	_ = h.userService.DeleteUser(ctx, user)

	responses.NewResponseJson(http.StatusNoContent).Write(w)
}
