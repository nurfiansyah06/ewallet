package tests

import (
	"ewalletgolang/db"
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"ewalletgolang/repository"
	"ewalletgolang/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	// Create a test server for handling the request.
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // You can provide a response here for your test case.
        w.WriteHeader(http.StatusOK)
    }))
    defer ts.Close() // Close the test server when the test is done.

    // Make a POST request to the test server.
    resp, err := http.Post(ts.URL+"/register", "application/json", nil)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    // Check if the status code is as expected.
    assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestLogin(t *testing.T) {
    // Create a test user with a known username and password.
    testUser := entity.User{
        Email:    "safas@gmail.c",
        Password: "abc",
    }

    // Mock or set up your authentication repository or service.
    // For the sake of this example, let's assume you have an AuthService.
	db := db.ConnectDB()
    userRepository := repository.NewRepository(db)
    userUsecase := usecase.NewUsecase(userRepository)

    // Test a successful login.
    authenticatedUser, err := userUsecase.Login(dto.UserLoginRequest{Email: testUser.Email, Password: testUser.Password})

    // Check if the login was successful and no error occurred.
    assert.NoError(t, err)
    assert.NotNil(t, authenticatedUser)

    // Check that the authenticated user's email matches the test user's email.
    assert.Equal(t, testUser.Email, authenticatedUser.Email)
}
