package endpoints

import (
	"log"
	"net/http"
	"os"

	"github.com/kataras/blocks"
)

var views *blocks.Blocks

func init() {
	views = blocks.New("./src/views")
	if err := views.Load(); err != nil {
		log.Fatal(err)
	}
}
func VerifyPassword(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("password") != os.Getenv("CIPHER_VIEW_PASS") {
		views.ExecuteTemplate(w, "error", "base", map[string]interface{}{
			"Error": "Invalid password",
		})
	}
	http.Redirect(w, r, "/home", http.StatusFound)
}
