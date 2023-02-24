package nlp_test

import (
	"github.com/magiconair/properties/assert"
	"nlp"
	"testing"
)

func TestTokenize(t *testing.T) {
	text := "Who's on first?"
	expected := []string{"who", "s", "on", "first"}
	tokens := nlp.Tokenize(text)

	assert.Equal(t, tokens, expected)
}
