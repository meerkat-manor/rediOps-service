package api


import (
	"net/http"
	"sync"
	"os"
	"io/ioutil"
	"log"
	"encoding/json"

	"github.com/labstack/echo/v4"
)


type Rediops struct {
	Lock   sync.Mutex
	ConfigFilename string
	DataFolder string
	ConfigFolder string
}

func NewRediops(configFolder string, configFilename string, dataFolder string) *Rediops {

	return &Rediops{
		ConfigFolder: configFolder,
		ConfigFilename: configFilename,
		DataFolder: dataFolder,
	}
}


// A common error payload returned
// when the response code is not 2xx
type ErrorModel struct {
	// Error description, that shuld be less technical
	// and more user orientated where possible
	Message    *string `json:"message,omitempty"`
	Resolution *string `json:"resolution,omitempty"`

	// Status code as a string
	Status *string `json:"status,omitempty"`

	// Numerical value of the status code
	StatusCode *int `json:"statusCode,omitempty"`

	// Technical information for the error.
	//
	// This must not contain sensitive information
	Technical *string `json:"technical,omitempty"`
}


func sendGeneralError(ctx echo.Context, code int, message string) error {

	var s = ""

	genErr := ErrorModel{
		Message: &message,
		Status:   &s,
		StatusCode: &code,
	}
	
	err := ctx.JSON(code, genErr)
	return err
}


func (ro *Rediops) GetHealth(ctx echo.Context) error {

	ro.Lock.Lock()
	defer ro.Lock.Unlock()

	var item = HealthResponse {
		Status: "OK",
		Message: "Service available",
	}

	return ctx.JSON(http.StatusOK, item)
}

// Fetch the DevOps information
// (GET /.well-known/devops)
func (ro *Rediops) GetWellKnownDevopsJson(ctx echo.Context) error {

	// Check if the corresponding file exists
	fileName := ro.DataFolder + "/devops.json"
	if _, err := os.Stat(fileName); err == nil {
		var item = DevopsModel {}

		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal("Error when opening file: ", err)
			return sendGeneralError(ctx, http.StatusInternalServerError ,
				"Internal server error")
		}
	 
		err = json.Unmarshal(content, &item)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
			return sendGeneralError(ctx, http.StatusInternalServerError ,
				"Internal server error")
		}

		return ctx.JSON(http.StatusOK, item)
	}

	
	return sendGeneralError(ctx, http.StatusNotFound,
		"devops.json not found")
}

// List matching captured resources
// (GET /devops/)
func (ro *Rediops) GetDevops(ctx echo.Context, params GetDevopsParams) error {
	

	var ilist = []DevopsbriefModel {}

	// Check if the corresponding folder exists
	if _, err := os.Stat(ro.DataFolder); err == nil {

		files, err := ioutil.ReadDir(ro.DataFolder)
		if err != nil {
			log.Fatal(err)
		}
	
		for _, file := range files {

			if (!file.IsDir() && (file.Name() != "devops.json")) {
				content, err := ioutil.ReadFile(ro.DataFolder + "/"+ file.Name())
				if err != nil {
					log.Printf("Error when opening file: %s", err)
					return sendGeneralError(ctx, http.StatusInternalServerError ,
						"Internal server error")
				}
			
				var item = DevopsModel {}
				err = json.Unmarshal(content, &item)
				if err != nil {
					log.Printf("Error during Unmarshal(): %s", err)
					return sendGeneralError(ctx, http.StatusInternalServerError ,
						"Internal server error")
				}

				var brief = DevopsbriefModel {
					Name: item.Name,
					Version: item.Version,
				}
				brief.UniqueId = *item.UniqueId
				brief.Guide = *item.Guide
				brief.Self = "/devops/"+brief.UniqueId
		
				ilist = append(ilist, brief)
			}
		}

		return ctx.JSON(http.StatusOK, ilist)
	} 
	
	return sendGeneralError(ctx, http.StatusNotFound,
		"Not found")


}

// (DELETE /devops/{id})
func (ro *Rediops) DeleteDevopsId(ctx echo.Context, id string) error {
	
	return sendGeneralError(ctx, http.StatusNotFound,
		"Not supported")

}

// Fetch the Devops resource
// (GET /devops/{id})
func (ro *Rediops) GetDevopsId(ctx echo.Context, id string) error {


	// Check if the corresponding file exists
	fileName := ro.DataFolder + "/"+ id + ".json"
	if _, err := os.Stat(fileName); err == nil {
		var item = DevopsModel {}

		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Printf("Error when opening file: %s", err)
			return sendGeneralError(ctx, http.StatusInternalServerError ,
				"Internal server error")
		}
	 
		err = json.Unmarshal(content, &item)
		if err != nil {
			log.Printf("Error during Unmarshal(): %s", err)
			return sendGeneralError(ctx, http.StatusInternalServerError ,
				"Internal server error")
		}

		return ctx.JSON(http.StatusOK, item)
	}

	return sendGeneralError(ctx, http.StatusNotFound,
		(id + " not found"))

}

// Update the DevOps resource
// (POST /devops/{id})
func (ro *Rediops) PostDevopsId(ctx echo.Context, id string) error {
	
	return sendGeneralError(ctx, http.StatusNotFound,
		"Not supported")

}
