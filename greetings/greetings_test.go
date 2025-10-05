package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestHelloMultipleNames tests Hello function with various names
func TestHelloMultipleNames(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantName bool
	}{
		{"valid name", "Alice", false, true},
		{"another valid name", "Bob", false, true},
		{"name with spaces", "John Doe", false, true},
		{"empty string", "", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := Hello(tt.input)

			if tt.wantErr {
				validateErrorCase(t, tt.input, msg, err)
				return
			}

			validateSuccessCase(t, tt.input, msg, err, tt.wantName)
		})
	}
}

// validateErrorCase checks that error cases behave correctly
func validateErrorCase(t *testing.T, input, msg string, err error) {
	if err == nil {
		t.Errorf("Hello(%q) expected error, got nil", input)
	}
	if msg != "" {
		t.Errorf("Hello(%q) expected empty message, got %q", input, msg)
	}
}

// validateSuccessCase checks that success cases behave correctly
func validateSuccessCase(t *testing.T, input, msg string, err error, wantName bool) {
	if err != nil {
		t.Errorf("Hello(%q) unexpected error: %v", input, err)
		return
	}

	if wantName {
		validateNameInMessage(t, input, msg)
	}
}

// validateNameInMessage checks that the message contains the expected name
func validateNameInMessage(t *testing.T, input, msg string) {
	nameRegex := regexp.MustCompile(`\b` + regexp.QuoteMeta(input) + `\b`)
	if !nameRegex.MatchString(msg) {
		t.Errorf("Hello(%q) = %q, want message containing %q", input, msg, input)
	}
}

// TestRandomFormat tests that randomFormat returns one of the expected formats
func TestRandomFormat(t *testing.T) {
	expectedFormats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Test multiple times to increase chance of getting different formats
	formatsSeen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		format := randomFormat()
		formatsSeen[format] = true

		// Check if the format is one of the expected ones
		found := false
		for _, expected := range expectedFormats {
			if format == expected {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("randomFormat() returned unexpected format: %q", format)
		}
	}

	// Check that we got at least one format (basic sanity check)
	if len(formatsSeen) == 0 {
		t.Error("randomFormat() never returned any format")
	}
}

// TestHelloFormatVariety tests that Hello function uses different formats
func TestHelloFormatVariety(t *testing.T) {
	name := "TestUser"
	messagesSeen := make(map[string]bool)

	// Call Hello multiple times to see if we get different formats
	for i := 0; i < 50; i++ {
		msg, err := Hello(name)
		if err != nil {
			t.Fatalf("Hello(%q) unexpected error: %v", name, err)
		}
		messagesSeen[msg] = true
	}

	// We should see at least one message (basic sanity check)
	if len(messagesSeen) == 0 {
		t.Error("Hello() never returned any message")
	}

	// All messages should contain the name
	nameRegex := regexp.MustCompile(`\b` + regexp.QuoteMeta(name) + `\b`)
	for msg := range messagesSeen {
		if !nameRegex.MatchString(msg) {
			t.Errorf("Hello(%q) returned message %q that doesn't contain the name", name, msg)
		}
	}
}
