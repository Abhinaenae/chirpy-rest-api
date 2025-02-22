package filter_test

import (
	"chirpy/internal/filter"
	"testing"
)

func TestFilterProfanity(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"This is a fuck example.", "This is a **** example."},
		{"Oh shit! That hurts.", "Oh ****! That hurts."},
		{"You are a bitch.", "You are a ****."},
		{"No profanity here.", "No profanity here."},
		{"What the fuck?", "What the ****?"},
		{"Shit happens.", "**** happens."},
		{"She is a bit chubby.", "She is a bit chubby."}, // Should not replace "bit"
		{"I said 'fuck' in quotes.", "I said '****' in quotes."},
		{"Fuck at the start.", "**** at the start."},
		{"Ends with shit", "Ends with ****"},
	}

	for _, tt := range tests {
		output := filter.FilterProfanity(tt.input)
		if output != tt.expected {
			t.Errorf("FilterProfanity(%q) = %q; want %q", tt.input, output, tt.expected)
		}
	}
}
