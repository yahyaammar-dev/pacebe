package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yahyaammar-dev/pacebe/configs"
	"github.com/yahyaammar-dev/pacebe/services/auth"
	email "github.com/yahyaammar-dev/pacebe/services/emails"
	"github.com/yahyaammar-dev/pacebe/services/event"
	"github.com/yahyaammar-dev/pacebe/types"
	"github.com/yahyaammar-dev/pacebe/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/reset-password", h.handleResetPassword).Methods("POST")
	router.HandleFunc("/create-password", h.handleCreatePassword).Methods("POST")
}

// @Summary Login
// @Description Logs in a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginPayload body types.LoginUserPayload true "Login user payload"
// @Success 200 {object} map[string]string "Successfully logged in"
// @Failure 400 {object} map[string]string "Invalid email or password"
// @Router /login [post]
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(configs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

// @Summary Register
// @Description Registers a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param registerPayload body types.RegisterUserPayload true "Register user payload"
// @Success 201 {object} map[string]string "User successfully registered"
// @Failure 400 {object} map[string]string "Invalid registration data"
// @Router /register [post]
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// check if user exists
	_, err := h.store.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	// hash password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	event.Dispatch(types.Event{
		Name:    "user.created",
		Payload: user.FirstName + " " + user.LastName,
	})

	utils.WriteJSON(w, http.StatusCreated, nil)
}

// @Summary Reset Password
// @Description Resets the user's password
// @Tags Auth
// @Accept json
// @Produce json
// @Param resetPasswordPayload body types.ResetPasswordPayload true "Reset password payload"
// @Success 200 {object} map[string]string "Password successfully reset"
// @Failure 400 {object} map[string]string "Invalid reset password data"
// @Router /reset-password [post]
func (h *Handler) handleResetPassword(w http.ResponseWriter, r *http.Request) {

	var resetPayload types.ResetPasswordPayload
	if err := utils.ParseJSON(r, &resetPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(resetPayload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	user, err := h.store.GetUserByEmail(resetPayload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	rndStr := utils.RandomString(15)

	err = h.store.UpdateUserRememberToken(user, rndStr)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	// send email to user
	email.SendEmail([]string{user.Email}, user.RememberToken)

	response := map[string]string{
		"status":  "success",
		"message": "Reset password link sent to email",
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

// @Summary Create Password
// @Description Creates a new password for the user
// @Tags Auth
// @Accept json
// @Produce json
// @Param createPasswordPayload body types.CreatePasswordPayload true "Create password payload"
// @Success 200 {object} map[string]string "Password successfully created"
// @Failure 400 {object} map[string]string "Invalid create password data"
// @Router /create-password [post]
func (h *Handler) handleCreatePassword(w http.ResponseWriter, r *http.Request) {
	var createPasswordPayload types.CreatePasswordPayload
	if err := utils.ParseJSON(r, &createPasswordPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(createPasswordPayload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	user, err := h.store.GetUserByRememberToken(createPasswordPayload.RememberToken)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	hashedPassword, err := auth.HashPassword(createPasswordPayload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.UpdatePasswordOfUser(user, hashedPassword)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Password Updated Successfully"})
}
