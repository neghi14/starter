package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/neghi14/starter/database"
	redisdb "github.com/neghi14/starter/database/plugins/redis"
	"github.com/neghi14/starter/utils"
)

func main() {

	r := chi.NewRouter()

	type UserModel struct {
		Email string `db:"email"`
		Name  string `db:"name"`
	}

	redisDB, err := redisdb.New(redisdb.Opts(), UserModel{})
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

	})

	panic(http.ListenAndServe(":8080", r))
}
