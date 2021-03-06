package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reedina/dryBuilder_services/ctrl"
	"github.com/reedina/dryBuilder_services/model"
	"github.com/rs/cors"

	//Initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//App  (TYPE)
type App struct {
	Router *mux.Router
}

//InitializeApplication - Init router, db connection and restful routes
func (a *App) InitializeApplication(user, password, url, dbname string) {

	model.ConnectDB(user, password, url, dbname)
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

//InitializeRoutes - Declare all application routes
func (a *App) InitializeRoutes() {

	//model.Region struct
	a.Router.HandleFunc("/api/aws/region", ctrl.CreateRegion).Methods("POST")
	a.Router.HandleFunc("/api/aws/regions", ctrl.GetRegions).Methods("GET")
	a.Router.HandleFunc("/api/aws/region/{id:[0-9]+}", ctrl.GetRegion).Methods("GET")
	a.Router.HandleFunc("/api/aws/region/{name}", ctrl.GetRegionByName).Methods("GET")
	a.Router.HandleFunc("/api/aws/region/{id:[0-9]+}", ctrl.UpdateRegion).Methods("PUT")
	a.Router.HandleFunc("/api/aws/region/{id:[0-9]+}", ctrl.DeleteRegion).Methods("DELETE")

	//model.InstanceType struct
	a.Router.HandleFunc("/api/aws/instance/type", ctrl.CreateInstanceType).Methods("POST")
	a.Router.HandleFunc("/api/aws/instance/types", ctrl.GetInstanceTypes).Methods("GET")
	a.Router.HandleFunc("/api/aws/instance/type/{id:[0-9]+}", ctrl.GetInstanceType).Methods("GET")
	a.Router.HandleFunc("/api/aws/instance/type/{name}", ctrl.GetInstanceTypeByName).Methods("GET")
	a.Router.HandleFunc("/api/aws/instance/type/{id:[0-9]+}", ctrl.UpdateInstanceType).Methods("PUT")
	a.Router.HandleFunc("/api/aws/instance/type/{id:[0-9]+}", ctrl.DeleteInstanceType).Methods("DELETE")

	//model.Auth struct
	a.Router.HandleFunc("/api/aws/auth", ctrl.CreateAuth).Methods("POST")
	a.Router.HandleFunc("/api/aws/auths", ctrl.GetAuths).Methods("GET")
	a.Router.HandleFunc("/api/aws/auth/{id:[0-9]+}", ctrl.GetAuth).Methods("GET")
	a.Router.HandleFunc("/api/aws/auth/{name}", ctrl.GetAuthByName).Methods("GET")
	a.Router.HandleFunc("/api/aws/auth/{id:[0-9]+}", ctrl.UpdateAuth).Methods("PUT")
	a.Router.HandleFunc("/api/aws/auth/{id:[0-9]+}", ctrl.DeleteAuth).Methods("DELETE")

	//model.PackerBuilderTYpe struct
	a.Router.HandleFunc("/api/packer/builder/type", ctrl.CreatePackerBuilderType).Methods("POST")
	a.Router.HandleFunc("/api/packer/builder/types", ctrl.GetPackerBuilderTypes).Methods("GET")
	a.Router.HandleFunc("/api/packer/builder/type/{id:[0-9]+}", ctrl.GetPackerBuilderType).Methods("GET")
	a.Router.HandleFunc("/api/packer/builder/type/{name}", ctrl.GetPackerBuilderTypeByName).Methods("GET")
	a.Router.HandleFunc("/api/packer/builder/type/{id:[0-9]+}", ctrl.UpdatePackerBuilderType).Methods("PUT")
	a.Router.HandleFunc("/api/packer/builder/type/{id:[0-9]+}", ctrl.DeletePackerBuilderType).Methods("DELETE")

	//model.AmiFilterLinux struct
	a.Router.HandleFunc("/api/packer/ami/filter/linux", ctrl.CreateAmiFilterLinux).Methods("POST")
	a.Router.HandleFunc("/api/packer/ami/filter/linuxes", ctrl.GetAmiFilterLinuxes).Methods("GET")
	a.Router.HandleFunc("/api/packer/ami/filter/linux/{id:[0-9]+}", ctrl.GetAmiFilterLinux).Methods("GET")
	a.Router.HandleFunc("/api/packer/ami/filter/linux/{name}", ctrl.GetAmiFilterLinuxByName).Methods("GET")
	a.Router.HandleFunc("/api/packer/ami/filter/linux/{id:[0-9]+}", ctrl.UpdateAmiFilterLinux).Methods("PUT")
	a.Router.HandleFunc("/api/packer/ami/filter/linux/{id:[0-9]+}", ctrl.DeleteAmiFilterLinux).Methods("DELETE")

	//model.AmiFilterLinux struct
	a.Router.HandleFunc("/api/packer/builder/ebs", ctrl.CreateEbsBuilder).Methods("POST")
	a.Router.HandleFunc("/api/packer/builder/ebses", ctrl.GetEbsBuilders).Methods("GET")
	a.Router.HandleFunc("/api/packer/builder/ebs/{id:[0-9]+}", ctrl.GetEbsBuilder).Methods("GET")
	a.Router.HandleFunc("/api/packer/builder/ebs/{name}", ctrl.GetEbsBuilderByName).Methods("GET")
	a.Router.HandleFunc("/api/packer/builder/ebs/{id:[0-9]+}", ctrl.UpdateEbsBuilder).Methods("PUT")
	a.Router.HandleFunc("/api/packer/builder/ebs/{id:[0-9]+}", ctrl.DeleteEbsBuilder).Methods("DELETE")

	//model.UserDataFile struct
	a.Router.HandleFunc("/api/packer/builder/userDataFile", ctrl.CreateUserDataFile).Methods("POST")
	a.Router.HandleFunc("/api/packer/builder/userDataFiles", ctrl.GetUserDataFiles).Methods("GET")
	a.Router.HandleFunc("/api/packer/builder/userDataFile/{id:[0-9]+}", ctrl.GetUserDataFile).Methods("GET")
	a.Router.HandleFunc("/api/packer/builder/userDataFile/{name}", ctrl.GetUserDataFileByName).Methods("GET")
	a.Router.HandleFunc("/api/packer/builder/userDataFile/{id:[0-9]+}", ctrl.UpdateUserDataFile).Methods("PUT")
	a.Router.HandleFunc("/api/packer/builder/userDataFile/{id:[0-9]+}", ctrl.DeleteUserDataFile).Methods("DELETE")

}

//RunApplication - Start the HTTP server
func (a *App) RunApplication(addr string) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	log.Fatal(http.ListenAndServe(addr, c.Handler(a.Router)))
}
