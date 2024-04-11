package task_manager_api

import "net/http"

func main() {
	ConnectDatabase()
	SetupRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error setting up server!")
	}
}
