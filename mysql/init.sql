CREATE DATABASE IF NOT EXISTS hackathon_backend;
USE hackathon_backend;
CREATE TABLE Clients (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    family_living_togethers VARCHAR(255)
);


CREATE TABLE States (
    id INT PRIMARY KEY AUTO_INCREMENT,
    disease VARCHAR(255) NOT NULL,
    treatments JSON NOT NULL,
    medicines JSON NOT NULL,
    treatment_policy VARCHAR(255) NOT NULL,
    client_id Int NOT NULL,
    FOREIGN KEY (client_id) REFERENCES Clients(id)
);

CREATE TABLE CarePlans (
    id INT PRIMARY KEY AUTO_INCREMENT,
    author VARCHAR(255),
    facility_name VARCHAR(255),
    result_analyze VARCHAR(255),
    care_committee_opinion VARCHAR(255),
    specified_service VARCHAR(255),
    care_policy VARCHAR(255),
    client_id Int NOT NULL,
    FOREIGN KEY (client_id) REFERENCES Clients(id),
    updated_at DATE
);


INSERT INTO Clients (name, age, family_living_togethers)
VALUES
    ('山田太郎', 65, '妻と二人暮らし'),
    ('鈴木花子', 45, '夫と子供2人と一緒に暮らしている'),
    ('田中一郎', 80, '一人暮らし');

INSERT INTO States (disease, treatments, medicines, treatment_policy, client_id)
VALUES
    ('高血圧',  '["足のリハビリ","点滴"]', '["薬A","薬B"]', '血圧の測定と投薬', 1),
    ('糖尿病', '["魁皇切開","ヘモグロビン輸入"]', '["薬B","薬C"]', '食事制限と運動療法', 2),
    ('肝炎', '["腕のマッサージ","緊急手術"]', '["薬A","薬C"]', '薬物療法と定期的な検査', 3);


INSERT INTO CarePlans (author, facility_name, result_analyze, care_committee_opinion, specified_service, care_policy, client_id, updated_at)
VALUES
    ('田中景子', '総合療養施設X', '検査結果による診断', '外来治療による経過観察を行うこと', '内服薬の投与と定期的な血液検査', '薬物療法の実施', 1, '2022-12-01'),
    ('山田花子', '総合療養施設X', '検査結果による診断', '外来治療による経過観察を行うこと', '内服薬の投与と定期的な血液検査', '薬物療法の実施', 1, '2022-12-01');

