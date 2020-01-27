package accesskey

import "github.com/rs/zerolog"

// MarshalZerologObject implements the zerolog marshaler so it can be logged using:
// log.With().EmbededObject(accesskey).Msg("Some message")
func (c Key) MarshalZerologObject(e *zerolog.Event) {
	e.Str("access_key", c.String())
}