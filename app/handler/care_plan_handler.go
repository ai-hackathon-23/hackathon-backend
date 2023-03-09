package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	rp "hackathon/repository"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		jsonData, _ := json.Marshal(CarePlan{
			Id:                   care_plan.Id,
			Author:               care_plan.Author.String,
			FacilityName:         care_plan.FacilityName.String,
			ResultAnalyze:        care_plan.ResultAnalyze.String,
			CareCommitteeOpinion: care_plan.CareCommitteeOpinion.String,
			SpecifiedService:     care_plan.SpecifiedService.String,
			CarePolicy:           care_plan.CarePolicy.String,
			UpdatedAt:            care_plan.UpdatedAt,
		})
		fmt.Fprintf(w, string(jsonData))
	}
	return nil
}

func (hd *CarePlanHandler) HandleGetCarePlans(w http.ResponseWriter, r *http.Request) error {
	care_plans, err := hd.rp.IndexCarePlan()
	if err != nil {
		log.Print(err)
		return err
	}

	cpList := []CarePlan{}
	for _, care_plan := range *care_plans {
		cp := CarePlan{
			Id:                   care_plan.Id,
			Author:               care_plan.Author.String,
			FacilityName:         care_plan.FacilityName.String,
			ResultAnalyze:        care_plan.ResultAnalyze.String,
			CareCommitteeOpinion: care_plan.CareCommitteeOpinion.String,
			SpecifiedService:     care_plan.SpecifiedService.String,
			CarePolicy:           care_plan.CarePolicy.String,
			UpdatedAt:            care_plan.UpdatedAt,
		}
		cpList = append(cpList, cp)
	}

	jsonData, err := json.Marshal(cpList)
	if err != nil {
		log.Print(err)
		return err
	}

	fmt.Fprintf(w, string(jsonData))
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
		UpdatedAt:            updatedAt,
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
			UpdatedAt:            updatedAt,
		})
		fmt.Fprintf(w, string(jsonData))
	}
	return nil
}

type CarePlan struct {
	Id                   int64  `json:"id"`
	Author               string `json:"author"`
	FacilityName         string `json:"facility_name"`
	ResultAnalyze        string `json:"result_analyze"`
	CareCommitteeOpinion string `json:"care_committee_opinion"`
	SpecifiedService     string `json:"specified_service"`
	CarePolicy           string `json:"care_policy"`
	UpdatedAt            string `json:"updated_at"`
}

func ToSnakeCase(s string) string {
	var builder strings.Builder
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				builder.WriteByte('_')
			}
			builder.WriteByte(byte(c + 32))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
