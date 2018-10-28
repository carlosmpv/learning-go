package index

import (
	"fmt"
	"net/http"

	"github.com/learning-go/golang-syntax-tests/middleware"
)

func Index(w http.ResponseWriter, r *http.Request, data []middleware.MidData) {
	fmt.Fprintf(w, "Index controller")
}
