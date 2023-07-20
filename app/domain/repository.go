package domain

import "context"

type PhrasesRepository interface {
	GetPhraseOfTheDay(ctx context.Context) (*Phrase, error)
}
