package config

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

func InitLogger(config *Config) {
	var writer io.Writer = os.Stdout

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	level, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if config.AppEnv == "local" {
		// Налаштовуємо ConsoleWriter
		writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,             // Писати в консоль
			TimeFormat: "2006-01-02 15:04:05", // Формат часу
			NoColor:    false,                 // Увімкнути кольори
		}

		defer func() {
			fmt.Println("--------------------------------------------------------------------")
			log.Trace().Msg("This is a trace message")
			log.Debug().Msg("This is a debug message")
			log.Info().Msg("This is an info message")
			log.Warn().Msg("This is a warning message")
			log.Error().Err(errors.New("some error")).Msg("This is an error message")
			fmt.Print("--------------------------------------------------------------------\n\n")
		}()
	}

	// Встановлюємо глобальний логер
	log.Logger = zerolog.New(writer).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(level)
}
