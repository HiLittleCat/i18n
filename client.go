package i18n

import (
	"golang.org/x/text/language"

	"github.com/HiLittleCat/core"
)

// ClientLocale returns the current locale used for the client.
func ClientLocale(c *core.Context) language.Tag {
	if v, ok := c.Data[contextDataKey]; ok {
		return v.(language.Tag)
	}
	return defaultLocale
}

// SetClientLocale sets the locale used for the client.
// If the locale described by language tag t doesn't exist, error ErrUnknownLocale is returned.
func SetClientLocale(c *core.Context, t language.Tag) error {
	if !locales.Has(t) {
		return ErrUnknownLocale
	}
	c.Data[contextDataKey] = t
	return nil
}
