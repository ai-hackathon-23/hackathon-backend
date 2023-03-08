package repository

import (
	"database/sql"
	"log"
	"time"
)

type CarePlanRepository struct {
	db *sql.DB
}

func NewCarePlanRepository(db *sql.DB) *CarePlanRepository {
	return &CarePlanRepository{db: db}
}

func (r *CarePlanRepository) CreateCarePlan(clientId string) (*CarePlan, error) {

	stmt, err := r.db.Prepare("INSERT INTO CarePlans(specified_service, care_policy, updated_at) VALUES(?,?,CURRENT_TIME)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec("歩行訓練や自立飲食ができるようにしていきましょう", "歌を口ずさむことに非常に生きがいを感じておられるので、喉元の治療はあまりしたくないそうです。そのため、喉を傷つけないよう、飲食介護の時には必ず職員が介助するようにします")
	if err != nil {
		return nil, err
	}
	lastId, _ := result.LastInsertId()

	stmt, err = r.db.Prepare("INSERT INTO CarePlanRecords(client_id,care_plan_id) VALUES(?,?)")
	if err != nil {
		return nil, err
	}
	result, err = stmt.Exec(clientId, lastId)
	t := time.Now().String()
	return &CarePlan{
		Id:               lastId,
		SpecifiedService: sql.NullString{String: "歩行訓練や自立飲食ができるようにしていきましょう", Valid: true},
		CarePolicy:       sql.NullString{String: "歌を口ずさむことに非常に生きがいを感じておられるので、喉元の治療はあまりしたくないそうです。そのため、喉を傷つけないよう、飲食介護の時には必ず職員が介助するようにします", Valid: true},
		UpdatedAt:        t,
	}, nil

}

func (r *CarePlanRepository) UpdateCarePlan(carePlan CarePlan) (*CarePlan, error) {

	// build the query string and the parameter list
	query := "UPDATE CarePlans SET"
	params := []interface{}{}
	if carePlan.Author.String != "" {
		query += " author = ?,"
		params = append(params, carePlan.Author)
	}
	if carePlan.FacilityName.String != "" {
		query += " facility_name = ?,"
		params = append(params, carePlan.FacilityName)
	}
	if carePlan.ResultAnalyze.String != "" {
		query += " result_analyze = ?,"
		params = append(params, carePlan.ResultAnalyze)
	}
	if carePlan.CareCommitteeOpinion.String != "" {
		query += " care_committee_opinion = ?,"
		params = append(params, carePlan.CareCommitteeOpinion)
	}
	if carePlan.SpecifiedService.String != "" {
		query += " specified_service = ?,"
		params = append(params, carePlan.SpecifiedService)
	}
	if carePlan.CarePolicy.String != "" {
		query += " care_policy = ?,"
		params = append(params, carePlan.CarePolicy)
	}
	// remove the trailing comma
	query = query[:len(query)-1]
	query += " WHERE id = ?"
	params = append(params, carePlan.Id)
	log.Print(query)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(params...)
	if err != nil {
		return nil, err
	}

	// 更新後のレコードを取得する
	updatedCarePlan, _ := r.GetCarePlanById(carePlan.Id)

	return updatedCarePlan, nil
}

func (r *CarePlanRepository) GetCarePlanById(id int64) (*CarePlan, error) {
	stmt, err := r.db.Prepare("SELECT * FROM CarePlans Where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	carePlan := CarePlan{}
	log.Print(row)
	err = row.Scan(
		&carePlan.Id,
		&carePlan.Author,
		&carePlan.FacilityName,
		&carePlan.ResultAnalyze,
		&carePlan.CareCommitteeOpinion,
		&carePlan.SpecifiedService,
		&carePlan.CarePolicy,
		&carePlan.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	log.Print(carePlan)

	return &carePlan, nil
}

type CarePlan struct {
	Id                   int64
	Author               sql.NullString
	FacilityName         sql.NullString
	ResultAnalyze        sql.NullString
	CareCommitteeOpinion sql.NullString
	SpecifiedService     sql.NullString
	CarePolicy           sql.NullString
	UpdatedAt            string
}
