package dashboard

import (
	"net/http"
	"github.com/archangel/SMx/config"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}



	config.TPL.ExecuteTemplate(w, "dashboard-index.html", nil)
}
