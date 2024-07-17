package model

type VatsimPilot struct {
	Cid        int        `json:"cid"`
	Name       string     `json:"name"`
	Callsign   string     `json:"callsign"`
	FlightPlan FlightPlan `json:"flight_plan"`
}

type VatprcPilot struct {
	Cid       int    `json:"cid"`
	Name      string `json:"name"`
	Callsign  string `json:"callsign"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	Aircraft  string `json:"aircraft"`
}

func isVatprcPilot(vatsimPilot *VatsimPilot) bool {
	if vatprcAirportRegexp.MatchString(vatsimPilot.FlightPlan.Departure) {
		return true
	}
	if vatprcAirportRegexp.MatchString(vatsimPilot.FlightPlan.Arrival) {
		return true
	}
	return false
}

func (vatsimPilot *VatsimPilot) MakeVatprcPilot() *VatprcPilot {
	if !isVatprcPilot(vatsimPilot) {
		return nil
	}
	var vatprcPilot VatprcPilot
	vatprcPilot.Cid = vatsimPilot.Cid
	vatprcPilot.Name = vatsimPilot.Name
	vatprcPilot.Callsign = vatsimPilot.Callsign
	vatprcPilot.Departure = vatsimPilot.FlightPlan.Departure
	vatprcPilot.Arrival = vatsimPilot.FlightPlan.Arrival
	vatprcPilot.Aircraft = vatsimPilot.FlightPlan.Aircraft
	return &vatprcPilot
}
