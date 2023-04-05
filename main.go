package main

import "auth_service/cmd"

func main() {
	r := cmd.InitializeRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
