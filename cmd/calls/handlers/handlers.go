package handlers

import (
	"CPaaS/internal/adapters/primary"
	"CPaaS/internal/adapters/secondary"
	cfg "CPaaS/internal/config"
	"fmt"
	"net/http"
)

func MakeCallHandler(w http.ResponseWriter, r *http.Request) {
	println("MakeCallHandler")
	from := r.FormValue("from")
	to := r.FormValue("to")
	actionUrl := r.FormValue("actionUrl")

	config := cfg.NewConfig([]string{from, to, actionUrl})
	secondaryPort := secondary.NewCallsApi(config) // The Secondary Adapter
	primaryPort := primary.NewService(secondaryPort)
	primaryAdapter := primary.NewHTTPPrimaryAdapter(primaryPort)

	result, err := primaryAdapter.HandleCall(cfg.MakeCall)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
