package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/kaiquecaires/hotel_reservation/db"
	"github.com/kaiquecaires/hotel_reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testDbUri  = "mongodb://localhost:27017"
	testDbName = "hotel-reservation-test"
)

type testDb struct {
	db.UserStore
}

func (tdb *testDb) tearDown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatalf(err.Error())
	}
}

func setup(t *testing.T) *testDb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDbUri))

	if err != nil {
		t.Fatalf(err.Error())
	}

	return &testDb{
		UserStore: db.NewMongoUserStore(client, testDbName),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "foo@bar.com",
		FirstName: "foo",
		LastName:  "bar",
		Password:  "123456789",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("post user expects %s, but received %s", "200", resp.Status)
		return
	}

	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	if user.FirstName != params.FirstName {
		t.Errorf("expected firstName %s, but got %s", params.FirstName, user.FirstName)
	}

	if user.LastName != params.LastName {
		t.Errorf("expected lastName %s, but got %s", params.LastName, user.LastName)
	}

	if user.Email != params.Email {
		t.Errorf("expected email %s, but got %s", params.Email, user.Email)
	}

	if user.ID == "" {
		t.Errorf("expected userID inserted")
	}

	if user.EncryptedPassowrd != "" {
		t.Errorf("expected EncryptedPassowrd not be included in the JSON response")
	}
}
