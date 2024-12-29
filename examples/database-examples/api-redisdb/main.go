package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/neghi14/starter/database"
	redisdb "github.com/neghi14/starter/database/plugins/redisdb"
	"github.com/neghi14/starter/utils"
)

func main() {

	r := chi.NewRouter()

	type UserModel struct {
		Email string `db:"email" json:"email"`
		Name  string `db:"name" json:"name"`
	}

	redisDB, err := redisdb.New(redisdb.Opts().SetConnectionUrl("localhost:6379").SetDatabase(0).SetTable("users"), UserModel{})
	if err != nil {
		panic(err)
	}

	r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		user, err := redisDB.FindOne(r.Context(), *database.Opts().Params(database.Param{Key: "id", Value: id}))
		if err != nil {
			utils.JSON(w).SetStatus(utils.ResponseError).SetStatusCode(http.StatusBadRequest).SetMessage(err.Error()).Send()
			return
		}
		utils.JSON(w).SetStatus(utils.ResponseSuccess).SetStatusCode(http.StatusOK).SetData(user).Send()
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		user, err := redisDB.Find(r.Context(), *database.Opts())
		if err != nil {
			utils.JSON(w).SetStatus(utils.ResponseError).SetStatusCode(http.StatusBadRequest).SetMessage(err.Error()).Send()
			return
		}
		utils.JSON(w).SetStatus(utils.ResponseSuccess).SetStatusCode(http.StatusOK).SetData(user).Send()
	})

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		var body UserModel

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			utils.JSON(w).SetStatus(utils.ResponseError).SetStatusCode(http.StatusBadRequest).SetMessage(err.Error()).Send()
			return
		}

		err = redisDB.Save(r.Context(), body)
		if err != nil {
			utils.JSON(w).SetStatus(utils.ResponseError).SetStatusCode(http.StatusBadRequest).SetMessage(err.Error()).Send()
			return
		}

			utils.JSON(w).SetStatus(utils.ResponseSuccess).SetStatusCode(http.StatusCreated).Send()
	})

	fmt.Println("Connection success")
	panic(http.ListenAndServe(":8080", r))
}
