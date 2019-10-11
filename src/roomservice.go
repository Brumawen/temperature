package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// RoomService contains service methods for the room being monitored
type RoomService struct {
	Srv *Server
}

// UpdateTelemetry will update all telemetry associated with the room
func (r *RoomService) UpdateTelemetry() error {
	// Update the last read time
	r.Srv.Room.LastRead = time.Now().UTC()

	out, err := exec.Command("python", "dht11.py").CombinedOutput()
	if err != nil {
		r.logError("Failed to read the DHT11 device", err.Error)
		return err
	}

	outStr := strings.TrimSpace(string(out))
	r.logDebug("Values returned =", outStr)
	dhtVals := strings.Split(outStr, ",")

	if f, err := strconv.ParseFloat(dhtVals[0], 64); err != nil {
		r.logError("Failed to get temperature value. " + err.Error() + ".")
	} else {
		r.Srv.Room.Temperature = f
	}

	if f, err := strconv.ParseFloat(dhtVals[1], 64); err != nil {
		r.logError("Failed to get humidity value. " + err.Error() + ".")
	} else {
		r.Srv.Room.Humidity = f
	}

	return nil
}

// logDebug logs a debug message to the logger
func (r *RoomService) logDebug(v ...interface{}) {
	if r.Srv.VerboseLogging {
		a := fmt.Sprint(v)
		logger.Info("RoomService: [Dbg] ", a[1:len(a)-1])
	}
}

// logInfo logs an information message to the logger
func (r *RoomService) logInfo(v ...interface{}) {
	a := fmt.Sprint(v)
	logger.Info("RoomService: [Inf] ", a[1:len(a)-1])
}

// logError logs an error message to the logger
func (r *RoomService) logError(v ...interface{}) {
	a := fmt.Sprint(v)
	logger.Error("RoomService [Err] ", a[1:len(a)-1])
}
