package repository

import (
	"database/sql"
)

type CarePlanRepository struct {
	db *sql.DB
}

func NewCarePlanRepository(db *sql.DB) *CarePlanRepository {
	return &CarePlanRepository{db: db}
}

func (r *CarePlanRepository) CreateCarePlan(clientId string) (*CarePlan, error) {

	stmt, err := r.db.Prepare("INSERT INTO CarePlans(specified_service,care_policy) VALUES(?,?)")
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
	return &CarePlan{Id: lastId, SpecifiedService: "歩行訓練や自立飲食ができるようにしていきましょう", CarePolicy: "歌を口ずさむことに非常に生きがいを感じておられるので、喉元の治療はあまりしたくないそうです。そのため、喉を傷つけないよう、飲食介護の時には必ず職員が介助するようにします"}, nil
}

type CarePlan struct {
	Id                 int64
	Author             string
	FacilityName       string
	ResultAnalyze      string
	CareComiteeOpinion string
	SpecifiedService   string
	CarePolicy         string
}
