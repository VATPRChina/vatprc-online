package model

type VatsimResponseGeneral struct {
	UpdateTimestamp string `json:"update_timestamp"`
}

type VatsimResponse struct {
	General     VatsimResponseGeneral `json:"general"`
	Pilots      []VatsimPilot         `json:"pilots"`
	Controllers []Controller          `json:"controllers"`
}

type VatprcStatus struct {
	LastUpdated       string             `json:"lastUpdated"`
	Pilots            []VatprcPilot      `json:"pilots"`
	Controllers       []Controller       `json:"controllers"`
	FutureControllers []FutureController `json:"futureControllers"`
}

func (vatsimResponse *VatsimResponse) MakeVatprcStatus() *VatprcStatus {
	var vatprcStatus VatprcStatus
	vatprcStatus.LastUpdated = vatsimResponse.General.UpdateTimestamp
	vatprcStatus.Pilots = make([]VatprcPilot, 0)
	for _, vatsimPilot := range vatsimResponse.Pilots {
		vatprcPilot := vatsimPilot.MakeVatprcPilot()
		if vatprcPilot != nil {
			vatprcStatus.Pilots = append(vatprcStatus.Pilots, *vatprcPilot)
		}
	}
	vatprcStatus.Controllers = make([]Controller, 0)
	for _, vatsimController := range vatsimResponse.Controllers {
		vatprcController := vatsimController.MakeVatprcController()
		if vatprcController != nil {
			vatprcStatus.Controllers = append(vatprcStatus.Controllers, *vatprcController)
		}
	}
	vatprcStatus.FutureControllers = make([]FutureController, 0)

	return &vatprcStatus
}
