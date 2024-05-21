package locator

import "testing"

// Verify that a new Locator instance is created with the correct base directory
func TestNewLocatorWithCorrectBaseDirectory(t *testing.T) {
	expectedDir := "/expected/directory"
	locator := NewLocator(expectedDir)
	if locator.BaseDir != expectedDir {
		t.Errorf("Expected base directory to be %s, got %s", expectedDir, locator.BaseDir)
	}
}

// test with an empty string to see how it handles it.
func TestNewLocatorWithEmptyString(t *testing.T) {
	locator := NewLocator("")
	if locator.BaseDir != "" {
		t.Errorf("Expected base directory to be an empty string, got %s", locator.BaseDir)
	}
}
