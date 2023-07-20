package repository

import (
	"bufio"
	"context"
	"encoding/csv"
	"github.com/Hudayberdyyev/log_with_trace/app/domain"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"strings"
)

type filePhraseRepository struct {
	storage *os.File
}

func NewFilePhraseRepository(storage *os.File) domain.PhrasesRepository {
	return &filePhraseRepository{storage: storage}
}

func (repo *filePhraseRepository) GetPhraseOfTheDay(_ context.Context) (*domain.Phrase, error) {
	fileScanner := bufio.NewScanner(repo.storage)
	fileScanner.Split(bufio.ScanLines)

	// Line scanning
	randomLineNumber := rand.Int31n(50)
	var lineNumber int32
	fileScanner.Scan()
	line := fileScanner.Text()
	log.Debug().Msgf("randomLineNumber = %d", randomLineNumber)
	for lineNumber = 1; lineNumber < randomLineNumber && fileScanner.Scan(); lineNumber++ {
		line = fileScanner.Text()
	}
	log.Debug().Msgf("line = %s, lineNumber = %d", line, lineNumber)

	// Separate phrase and author
	r := csv.NewReader(strings.NewReader(line))
	r.Comma = '\t'
	record, err := r.Read()
	if err != nil {
		log.Error().Err(err).
			Str("line", line).
			Int32("randomLineNumber", randomLineNumber).
			Int32("lineNumber", lineNumber).
			Msg("error reading csv line")
		return nil, errors.New(err.Error())
	}
	return domain.NewPhrase(record[0], record[1]), nil
}
