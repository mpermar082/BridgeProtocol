// internal/bridgeprotocol/bridgeprotocol_test.go
package bridgeprotocol

import (
    "testing"
)

// TestNewApp verifies that NewApp returns a non-nil app with expected properties.
func TestNewApp(t *testing.T) {
    // Create a new app with verbose logging enabled
    app := NewApp(true)
    if app == nil {
        t.Fatal("NewApp returned nil")
    }
    if !app.Verbose {
        t.Errorf("Expected verbose to be true, got %v", app.Verbose)
    }
    if app.ProcessedCount != 0 {
        t.Errorf("Expected ProcessedCount to be 0, got %d", app.ProcessedCount)
    }
}

// TestProcess verifies that the Process method returns a successful result.
func TestProcess(t *testing.T) {
    // Create a new app with verbose logging disabled
    app := NewApp(false)
    result, err := app.Process("test data")
    
    if err != nil {
        t.Fatalf("Process returned error: %v", err)
    }
    
    if !result.Success {
        t.Errorf("Expected result.Success to be true, got %v", result.Success)
    }
    
    if app.ProcessedCount != 1 {
        t.Errorf("Expected ProcessedCount to be 1, got %d", app.ProcessedCount)
    }
}

// TestRun verifies that the Run method returns no error.
func TestRun(t *testing.T) {
    // Create a new app with verbose logging disabled
    app := NewApp(false)
    err := app.Run("", "")
    
    if err != nil {
        t.Errorf("Run returned unexpected error: %v", err)
    }
}