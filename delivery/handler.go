package delivery 

import (
	"net/http"

  "portifolio/pages"
	"github.com/angelofallars/htmx-go"
)

func HandlePortifolio() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := htmx.NewResponse().Refresh(true)
		writer.RenderTempl(r.Context(), w, pages.Portifolio("Will"))
		writer.Write(w)
	}
}
