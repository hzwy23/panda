package route

import (
	"fmt"
	"net/http"
)

type ControllerHandle interface{
	Get(w http.ResponseWriter,r *http.Request)
	Post(w http.ResponseWriter,r *http.Request)
	Delete(w http.ResponseWriter,r *http.Request)
	Put(w http.ResponseWriter,r *http.Request)
}

type Controller struct{
}

func (t *Controller)Get(w http.ResponseWriter,r *http.Request){
	fmt.Println("default implement")
}

func (t *Controller)Post(w http.ResponseWriter,r *http.Request){
	fmt.Println("default implement")
}

func (t *Controller)Delete(w http.ResponseWriter,r *http.Request){
	fmt.Println("default implement")
}

func (t *Controller)Put(w http.ResponseWriter,r *http.Request){
	fmt.Println("default implement")
}