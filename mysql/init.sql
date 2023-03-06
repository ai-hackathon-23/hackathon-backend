CREATE DATABASE IF NOT EXISTS hackathon_backend;
USE hackathon_backend;

CREATE TABLE States (
    id INT PRIMARY KEY AUTO_INCREMENT,
    disease VARCHAR(255) NOT NULL,
    treatments JSON NOT NULL,
    medicines JSON NOT NULL,
    treatment_policy VARCHAR(255) NOT NULL
);

CREATE TABLE CarePlans (
    id INT PRIMARY KEY AUTO_INCREMENT,
    author VARCHAR(255) NOT NULL,
    facility_name VARCHAR(255) NOT NULL,
    result_analyze VARCHAR(255),
    care_committee_opinion VARCHAR(255),
    specified_service VARCHAR(255),
    care_policy VARCHAR(255)
);

CREATE TABLE Clients (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    family_living_together JSON
);

CREATE TABLE StateRecords (
state_id INT NOT NULL,
client_id Int NOT NULL,
FOREIGN KEY (state_id) REFERENCES States(id),
FOREIGN KEY (client_id) REFERENCES Clients(id)
);

CREATE TABLE CarePlanRecords (
care_plan_id INT,
client_id Int,
FOREIGN KEY (care_plan_id) REFERENCES CarePlans(id),
FOREIGN KEY (client_id) REFERENCES Clients(id)
);

