package update

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

type Release struct {
	Tag string `json:"tag_name"`
}

// Checks the current release of the CLI and informs the user

func CheckReleases(currentVersion string) {
	resp, err := http.Get("https://api.github.com/repos/Platon223/Larb/releases/latest")
	if err != nil {
		return // not critical, so silent return
	}

	defer resp.Body.Close()

	var currentRelease Release
	json.NewDecoder(resp.Body).Decode(&currentRelease)

	if currentRelease.Tag != currentVersion {
		updateCmd := ""
		switch runtime.GOOS {
		case "darwin":
			updateCmd = "brew upgrade --cask Platon223/tap/larb"
		case "linux":
			updateCmd = "curl -sSL https://github.com/Platon223/Larb/releases/latest/download/larb_linux_amd64.tar.gz | tar -xz && sudo mv larb /usr/local/bin/"
		case "windows":
			updateCmd = "Download the latest release from https://github.com/Platon223/Larb/releases/latest"
		}

		fmt.Println()
		fmt.Printf("\n⚠️  A new version of LARB is available: %s → %s\n", currentVersion, currentRelease.Tag)
		fmt.Printf("   Update with: %s", updateCmd)
		fmt.Println()
	}
}
