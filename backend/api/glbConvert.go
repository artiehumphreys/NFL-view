package api

import (
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func GLBConvert() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		gameID := ps.ByName("gameId")
		playID := ps.ByName("playId")

		basePath, err := filepath.Abs(filepath.Dir("."))
		if err != nil {
			http.Error(w, "Failed to determine base path: "+err.Error(), http.StatusInternalServerError)
			return
		}
		scriptPath := filepath.Join(basePath, "pkg", "bvh_to_glb.py")
		requirementsPath := filepath.Join(basePath, "requirements.txt")

		pipInstallCmd := exec.Command("python", "-m", "pip", "install", "-r", requirementsPath)
		output, err := pipInstallCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to install Python libraries: %s\n%s", err.Error(), string(output))
			http.Error(w, "Failed to install Python libraries: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Python libraries installed successfully")

		cmd := exec.Command("python", scriptPath, gameID, playID)
		cmdOutput, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to execute Python script: %s\n%s", err.Error(), string(cmdOutput))
			http.Error(w, "Failed to execute Python script: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Python script executed successfully")
		fmt.Println(string(cmdOutput))

		w.WriteHeader(http.StatusNoContent)
	}
}
