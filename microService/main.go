package microService

import (
	"GoPlayground/microService/controller"
	"net/http"
)

func main() {
	dbUrl := "mongodb://localhost:27017"
	mux := http.NewServeMux()
	_ = controller.StartControllers(dbUrl, mux)
	http.ListenAndServe(":8080", mux)
}
