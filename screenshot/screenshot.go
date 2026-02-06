package screenshot

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// CaptureScreen captures a screenshot and saves it to a temporary file
// Returns the path to the saved screenshot
func CaptureScreen() (string, error) {
	// Create temp directory if it doesn't exist
	tempDir := filepath.Join(os.TempDir(), "textgrab")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}

	// Generate unique filename
	timestamp := time.Now().Format("20060102_150405")
	filename := filepath.Join(tempDir, fmt.Sprintf("screenshot_%s.png", timestamp))

	// Use macOS screencapture command with interactive selection
	// -i: interactive mode (user selects area)
	// -s: selection mode only (no window shadow)
	cmd := exec.Command("screencapture", "-i", "-s", filename)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to capture screenshot: %w", err)
	}

	// Check if file was created (user might have cancelled)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", fmt.Errorf("screenshot cancelled by user")
	}

	return filename, nil
}

// CaptureFullScreen captures the entire screen without user interaction
func CaptureFullScreen() (string, error) {
	tempDir := filepath.Join(os.TempDir(), "textgrab")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}

	timestamp := time.Now().Format("20060102_150405")
	filename := filepath.Join(tempDir, fmt.Sprintf("screenshot_%s.png", timestamp))

	// Capture full screen without interaction
	cmd := exec.Command("screencapture", filename)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to capture screenshot: %w", err)
	}

	return filename, nil
}

// CaptureRegion captures a specific region of the screen
// x, y: top-left corner coordinates
// width, height: dimensions of the region
func CaptureRegion(x, y, width, height int) (string, error) {
	tempDir := filepath.Join(os.TempDir(), "textgrab")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}

	timestamp := time.Now().Format("20060102_150405")
	filename := filepath.Join(tempDir, fmt.Sprintf("screenshot_%s.png", timestamp))

	// -R: capture specific region (x,y,width,height)
	region := fmt.Sprintf("%d,%d,%d,%d", x, y, width, height)
	cmd := exec.Command("screencapture", "-R", region, filename)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to capture screenshot: %w", err)
	}

	return filename, nil
}

// Cleanup removes the temporary screenshot file
func Cleanup(filepath string) error {
	return os.Remove(filepath)
}
