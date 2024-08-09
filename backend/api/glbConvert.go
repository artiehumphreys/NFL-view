package api

import (
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func GLBConvert() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		gameID := ps.ByName("gameId")
		playID := ps.ByName("playId")
		cmd := exec.Command("python3", "/absolute/path/to/pkg/bvh_to_glb.py", gameID, playID)

		err := cmd.Run()
		if err != nil {
			http.Error(w, "Failed to execute script: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
