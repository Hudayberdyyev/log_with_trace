package app

import (
	"github.com/Hudayberdyyev/log_with_trace/pkg/handler"
	"github.com/Hudayberdyyev/log_with_trace/pkg/repository"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func Init() {
	file, err := os.Open("data/phrases.data")
	if err != nil {
		panic(err)
	}

	repo := repository.NewFilePhraseRepository(file)
	handlers := handler.NewHandler(repo)

	log.Info().Msg("Starting HTTP server")
	if err := http.ListenAndServe(":3333", handlers); err != nil {
		log.Fatal().Err(err)
	}
	log.Info().Msg("Exit HTTP server")
}
