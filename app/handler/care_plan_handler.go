package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	rp "hackathon/repository"
	"log"
	"net/http"
	"strconv"
)

type CarePlanHandler struct {
	rp *rp.CarePlanRepository
}

func NewCarePlanHandler(repository *rp.CarePlanRepository) CarePlanHandler {
	return CarePlanHandler{repository}
}

func (hd *CarePlanHandler) HandleGetCarePlan(w http.ResponseWriter, r *http.Request) error {
	clientId := r.FormValue("client_id")

	care_plan, err := hd.rp.CreateCarePlan(clientId)
	if err != nil {
		log.Print(err)
	} else {
		jsonData, _ := json.Marshal(care_plan)
		fmt.Fprintf(w, string(jsonData))
	}
	return nil
}

func (hd *CarePlanHandler) HandleUpdateCarePlan(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	author := r.FormValue("author")
	facilityName := r.FormValue("facility_name")
	resultAnalyze := r.FormValue("result_analyze")
	careComitteeOpinion := r.FormValue("care_committee_opinion")
	specifiedService := r.FormValue("specified_service")
	carePolicy := r.FormValue("care_policy")
	updatedAt := r.FormValue("updated_at")
	carePlan := rp.CarePlan{
		Id:                   id,
		Author:               sql.NullString{String: author, Valid: author != ""},
		FacilityName:         sql.NullString{String: facilityName, Valid: facilityName != ""},
		ResultAnalyze:        sql.NullString{String: resultAnalyze, Valid: resultAnalyze != ""},
		CareCommitteeOpinion: sql.NullString{String: careComitteeOpinion, Valid: careComitteeOpinion != ""},
		SpecifiedService:     sql.NullString{String: specifiedService, Valid: specifiedService != ""},
		CarePolicy:           sql.NullString{String: carePolicy, Valid: carePolicy != ""},
		UpdatedAt:            sql.NullString{String: updatedAt, Valid: updatedAt != ""},
	}

	updatedCarePlan, err := hd.rp.UpdateCarePlan(carePlan)
	if err != nil {
		log.Print(err)
	} else {
		jsonData, _ := json.Marshal(CarePlan{
			Id:                   updatedCarePlan.Id,
			Author:               updatedCarePlan.Author.String,
			FacilityName:         updatedCarePlan.FacilityName.String,
			ResultAnalyze:        updatedCarePlan.ResultAnalyze.String,
			CareCommitteeOpinion: updatedCarePlan.CareCommitteeOpinion.String,
			SpecifiedService:     updatedCarePlan.SpecifiedService.String,
			CarePolicy:           updatedCarePlan.CarePolicy.String,
			UpdatedAt:            updatedCarePlan.UpdatedAt.String,
		})
		fmt.Fprintf(w, string(jsonData))
	}
	return nil
}

type CarePlan struct {
	Id                   int64
	Author               string
	FacilityName         string
	ResultAnalyze        string
	CareCommitteeOpinion string
	SpecifiedService     string
	CarePolicy           string
	UpdatedAt            string
}
