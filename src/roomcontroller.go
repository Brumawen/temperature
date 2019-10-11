package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RoomController handles the Web Methods for the room being monitored
type RoomController struct {
	Srv *Server
}

// AddController adds the controller routes to the router
func (c *RoomController) AddController(router *mux.Router, s *Server) {
	c.Srv = s
	router.Methods("GET").Path("/room/get").Name("GetTelemetry").
		Handler(Logger(c, http.HandlerFunc(c.handleGetTelemetry)))
}

// handlerGetTelemetry will return the current telemetry for the room
func (c *RoomController) handleGetTelemetry(w http.ResponseWriter, r *http.Request) {
	if err := c.Srv.RoomService.UpdateTelemetry(); err != nil {
		http.Error(w, "Error updating telemetry", http.StatusInternalServerError)
	} else {
		if err := c.Srv.Room.WriteTo(w); err != nil {
			c.LogError("Error serializing telemetry.", err.Error)
			http.Error(w, "Error serializing telemetry", http.StatusInternalServerError)
		}
	}
}

// LogInfo is used to log information messages for this controller.
func (c *RoomController) LogInfo(v ...interface{}) {
	a := fmt.Sprint(v)
	logger.Info("RoomController: [Inf] ", a[1:len(a)-1])
}

// LogError is used to log information messages for this controller.
func (c *RoomController) LogError(v ...interface{}) {
	a := fmt.Sprint(v)
	logger.Info("RoomController: [Err] ", a[1:len(a)-1])
}
