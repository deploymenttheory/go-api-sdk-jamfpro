package dependancies

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// downloadFile downloads a file from the specified URL and saves it to the specified local file path.
func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// installSuspiciousPackage downloads and installs Suspicious Package from the given URL.
func installSuspiciousPackage(downloadURL string) error {
	dmgFilePath := "/tmp/SuspiciousPackage.dmg"

	fmt.Println("Downloading Suspicious Package...")
	err := downloadFile(downloadURL, dmgFilePath)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	fmt.Println("Download completed.")

	fmt.Println("Mounting .dmg file...")
	mountCmd := exec.Command("hdiutil", "attach", dmgFilePath)
	if err := mountCmd.Run(); err != nil {
		return fmt.Errorf("failed to mount .dmg file: %w", err)
	}

	fmt.Println("Copying Suspicious Package to /Applications...")
	copyCmd := exec.Command("cp", "-r", "/Volumes/Suspicious Package/Suspicious Package.app", "/Applications/")
	if err := copyCmd.Run(); err != nil {
		return fmt.Errorf("failed to copy Suspicious Package to /Applications: %w", err)
	}

	fmt.Println("Installation completed. Please follow any additional on-screen instructions if necessary.")

	return nil
}
