package model

import (
	"fmt"
	"time"
)

type Controller struct {
	Cid       int    `json:"cid"`
	Name      string `json:"name"`
	Callsign  string `json:"callsign"`
	Frequency string `json:"frequency"`
}

type AtcCenterFutureController struct {
	Callsign string        `json:"callsign"`
	Start    string        `json:"start"`
	Finish   string        `json:"finish"`
	User     AtcCenterUser `json:"user"`
}

type AtcCenterUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type FutureController struct {
	Callsign string `json:"callsign"`
	Name     string `json:"name"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

func (controller *Controller) MakeVatprcController() *Controller {
	if !vatprcControllerRegexp.MatchString(controller.Callsign) {
		return nil
	}
	return &(*controller)
}

func (a *AtcCenterFutureController) MakeFutureController() *FutureController {
	if !vatprcControllerRegexp.MatchString(a.Callsign) {
		return nil
	}
	var result FutureController
	result.Name = fmt.Sprintf("%s %s", a.User.FirstName, a.User.LastName)
	result.Callsign = a.Callsign
	t, err := time.Parse("2006-01-02T15:04:05Z", a.Start)
	if err != nil {
		result.Start = ""
	} else {
		result.Start = t.Format("02 15:04")
	}
	t, err = time.Parse("2006-01-02T15:04:05Z", a.Finish)
	if err != nil {
		result.End = ""
	} else {
		result.End = t.Format("02 15:04")
	}
	return &result
}
