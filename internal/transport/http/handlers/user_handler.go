package handlers

import (
	"github.com/Andrew-UA/product-list/app/services"
	"github.com/Andrew-UA/product-list/internal/config"
	"github.com/Andrew-UA/product-list/internal/transport/http/requests/user_requests"
	"github.com/Andrew-UA/product-list/internal/transport/http/responses"
	"github.com/Andrew-UA/product-list/internal/transport/http/responses/user_responses"
	"net/http"
	"strconv"
)

type UserHandler struct {
	config      *config.Config
	userService services.UsersServiceInterface
}

func NewUserHandler(config *config.Config, userService services.UsersServiceInterface) *UserHandler {
	return &UserHandler{
		config:      config,
		userService: userService,
	}
}

func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	var response *user_responses.UserListResponse

	ctx := r.Context()
	users, _ := h.userService.GetUsers(ctx)

	response = user_responses.NewUserListResponse(users)
	response.SetStatusCode(http.StatusOK)
	response.Write(w)
}

func (h *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	var response *user_responses.UserResponse

	ctx := r.Context()

	userIdStr := r.URL.Query().Get("id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	user, _ := h.userService.GetUserById(ctx, userId)

	response = user_responses.NewUserResponse(user)
	response.SetStatusCode(http.StatusOK)
	response.Write(w)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request user_requests.CreateUserRequest
	var response *user_responses.UserResponse

	ctx := r.Context()
	_ = request.ReadAndClose(r.Body)
	_ = request.Validate()

	data := request.ToMap()
	user, _ := h.userService.CreateUser(ctx, data)

	response = user_responses.NewUserResponse(user)
	response.SetStatusCode(http.StatusCreated)
	response.Write(w)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var request user_requests.UpdateUserRequest
	var response *user_responses.UserResponse

	ctx := r.Context()
	userIdStr := r.URL.Query().Get("id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	user, _ := h.userService.GetUserById(ctx, userId)

	_ = request.ReadAndClose(r.Body)
	_ = request.Validate()

	data := request.ToMap()
	user, _ = h.userService.UpdateUser(ctx, data, user)

	response = user_responses.NewUserResponse(user)
	response.SetStatusCode(http.StatusAccepted)
	response.Write(w)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var response *responses.Response

	ctx := r.Context()
	userIdStr := r.URL.Query().Get("id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	user, _ := h.userService.GetUserById(ctx, userId)

	_ = h.userService.DeleteUser(ctx, user)

	response = &responses.Response{}
	response.SetStatusCode(http.StatusAccepted)
	response.Write(w)
}
