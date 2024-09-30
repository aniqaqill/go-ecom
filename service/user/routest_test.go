package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aniqaqill/go-ecom/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "invalid",
			Password:  "password",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rr.Code)
		}
	})

	t.Run("should correctly register a user", func(t *testing.T) {
		userStore := &mockUserStore{}
		handler := NewHandler(userStore)

		t.Run("should fail if user payload is invalid", func(t *testing.T) {
			payload := types.RegisterUserPayload{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "valid@gmail.com",
				Password:  "password",
			}

			marshalled, _ := json.Marshal(payload)

			req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			router.HandleFunc("/register", handler.handleRegister)

			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusCreated {
				t.Errorf("expected status code 201, got %d", rr.Code)
			}
		})

	})

}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not founf")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}

func (m *mockUserStore) UpdateUser(u types.User) error {
	return nil
}

func (m *mockUserStore) DeleteUser(id int) error {
	return nil
}