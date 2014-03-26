/* Demonstrate basic CRUD operation using a store based on MySQL and GORM

GORM (github.com/jinzhu/gorm) is simple ORM library.
In this example the same struct is used both as the GORM model and as the JSON model.

The curl demo:

        curl -i -d '{"Message":"this is a test"}' http://127.0.0.1:8080/reminders
        curl -i http://127.0.0.1:8080/reminders/1
        curl -i http://127.0.0.1:8080/reminders
        curl -i -X PUT -d '{"Message":"is updated"}' http://127.0.0.1:8080/reminders/1
        curl -i -X DELETE http://127.0.0.1:8080/reminders/1

*/
package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
)

func main() {

	api := Api{}
	api.InitDB()
	api.InitSchema()

	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	handler.SetRoutes(
		rest.RouteObjectMethod("GET", "/reminders", &api, "GetAllReminders"),
		rest.RouteObjectMethod("POST", "/reminders", &api, "PostReminder"),
		rest.RouteObjectMethod("GET", "/reminders/:id", &api, "GetReminder"),
		rest.RouteObjectMethod("PUT", "/reminders/:id", &api, "PutReminder"),
		rest.RouteObjectMethod("DELETE", "/reminders/:id", &api, "DeleteReminder"),
	)
	http.ListenAndServe(":8080", &handler)
}

type Reminder struct {
	Id        int64     `json:"id"`
	Message   string    `sql:"size:1024" json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}

type Api struct {
	DB gorm.DB
}

func (api *Api) InitDB() {
	var err error
	api.DB, err = gorm.Open("mysql", "gorm:gorm@/gorm?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	api.DB.LogMode(true)
}

func (api *Api) InitSchema() {
	api.DB.AutoMigrate(Reminder{})
}

func (api *Api) GetAllReminders(w rest.ResponseWriter, r *rest.Request) {
	reminders := []Reminder{}
	api.DB.Find(&reminders)
	w.WriteJson(&reminders)
}

func (api *Api) GetReminder(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	reminder := Reminder{}
	if api.DB.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&reminder)
}

func (api *Api) PostReminder(w rest.ResponseWriter, r *rest.Request) {
	reminder := Reminder{}
	if err := r.DecodeJsonPayload(&reminder); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := api.DB.Save(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&reminder)
}

func (api *Api) PutReminder(w rest.ResponseWriter, r *rest.Request) {

	id := r.PathParam("id")
	reminder := Reminder{}
	if api.DB.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Reminder{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reminder.Message = updated.Message

	if err := api.DB.Save(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&reminder)
}

func (api *Api) DeleteReminder(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	reminder := Reminder{}
	if api.DB.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := api.DB.Delete(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
