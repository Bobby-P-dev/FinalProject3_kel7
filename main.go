package main

import (
	"github.com/Bobby-P-dev/FinalProject3_kel7/database"
	"github.com/Bobby-P-dev/FinalProject3_kel7/initiallizers"
	"github.com/Bobby-P-dev/FinalProject3_kel7/routers"
)

func init() {
	initiallizers.LoadEnvVariable()
}
func main() {
	database.ConnectToDB()
	r := routers.StarApp()
	r.Run()
}
