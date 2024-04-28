package main

import (
	"fmt"
	"net/http"

	"github.com/Tigraqt/rss/internal/database"
)

func (cfg *apiConfig) handlerGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get posts: %v", err))
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
