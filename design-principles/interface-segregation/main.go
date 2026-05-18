package main

import "fmt"

type FleetManager interface {
	GetLocation() (string, error)
}
type StreamManager interface {
	StreamVideo() ([]byte, error)
}
type BatteryManager interface {
	GetBatteryLevel() (int, error)
}
type HornManager interface {
	PlayHorn() error
}
type DiagnosticManager interface {
	RunDiagnostics() error
}

type GPSTracker struct{}

func (g *GPSTracker) GetLocation() (string, error) {
	return "lat:3.1390, lng:101.6869", nil
}

// Full self-driving unit — supports everything
type AutopilotUnit struct{}

func (a *AutopilotUnit) GetLocation() (string, error)  { return "lat:3.1390, lng:101.6869", nil }
func (a *AutopilotUnit) StreamVideo() ([]byte, error)  { return []byte("video-feed"), nil }
func (a *AutopilotUnit) GetBatteryLevel() (int, error) { return 87, nil }
func (a *AutopilotUnit) PlayHorn() error               { fmt.Println("Beep!"); return nil }
func (a *AutopilotUnit) RunDiagnostics() error         { fmt.Println("All systems go"); return nil }

// Each function asks for only what it needs
func trackLocation(f FleetManager) {
	loc, _ := f.GetLocation()
	fmt.Println("Location:", loc)
}

func streamFeed(s StreamManager) {
	data, _ := s.StreamVideo()
	fmt.Println("Streaming:", len(data), "bytes")
}

func runDiagnostics(d DiagnosticManager) {
	d.RunDiagnostics()
}

func playHorn(h HornManager) {
	h.PlayHorn()
}

func main() {
	gps := &GPSTracker{}
	autopilot := &AutopilotUnit{}

	// GPSTracker — only promised GetLocation, compiler enforces it
	trackLocation(gps)
	trackLocation(autopilot) // also works — AutopilotUnit satisfies FleetManager too

	// Only AutopilotUnit can do these — GPSTracker can't even enter
	streamFeed(autopilot)
	runDiagnostics(autopilot)
	playHorn(autopilot)
}
