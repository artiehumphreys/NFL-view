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

		exec.Command("python", "-m", "venv", "venv")
		exec.Command("source", "venv/bin/activate")
		cmd := exec.Command("python", "-m", "pip", "install", "-r", requirementsPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to install Python libraries: %s\n%s", err.Error(), string(output))
		}
		fmt.Println("Python libraries installed successfully")

		exec.Command("python", scriptPath, gameID, playID)
		fmt.Println("hi")
		cmd.Run()
		fmt.Println("hio")

		w.WriteHeader(http.StatusNoContent)
	}
}
