package main

import (
	"CPaaS/internal/adapters/primary"
	"CPaaS/internal/adapters/secondary"
	"CPaaS/internal/config"
	"fmt"
	"os"
)

func main() {

	c := config.NewConfig(os.Args[1:])
	secondaryPort := secondary.NewCallsApi(c) // The Secondary Adapter
	primaryPort := primary.NewService(secondaryPort)
	primaryAdapter := primary.NewCLIPrimaryAdapter(primaryPort)

	result, err := primaryAdapter.HandleCall(config.MakeCall)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
