package main

import (
	"github.com/Hudayberdyyev/log_with_trace/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	// Setup logger
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.With().Caller().Timestamp().Logger()

	app.Init()
}
