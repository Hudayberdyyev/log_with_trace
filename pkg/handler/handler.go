package handler

import (
	"bytes"
	"fmt"
	"github.com/Hudayberdyyev/log_with_trace/app/domain"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Handler struct {
	PhrasesRepo domain.PhrasesRepository
}

func NewHandler(phraseRepo domain.PhrasesRepository) *Handler {
	return &Handler{PhrasesRepo: phraseRepo}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/phrase-of-the-day":
		h.phraseOfTheDay(w, r)
	case "/greet":
		h.greet(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) phraseOfTheDay(w http.ResponseWriter, r *http.Request) {
	phrase, err := h.PhrasesRepo.GetPhraseOfTheDay(r.Context())
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		http.Error(w, "Произошла ошибка при обработке запроса. Попробуйте еще раз", http.StatusInternalServerError)
		return
	}

	buff := bytes.NewBufferString(fmt.Sprintf("%s\n%s", phrase.GetPhrase(), phrase.GetAuthor()))
	if _, err = w.Write(buff.Bytes()); err != nil {
		log.Error().Err(err).Msg("")
		http.Error(w, "Произошла ошибка при записи тела ответа. Попробуйте еще раз", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) greet(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	if _, err := w.Write([]byte("Greetings from server ...")); err != nil {
		log.Error().Err(err).Msg("")
		http.Error(w, "Произошла ошибка при записи тела ответа. Попробуйте еще раз", http.StatusInternalServerError)
		return
	}
}
