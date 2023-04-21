create table d_medical_organization
(
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    code varchar(50) null,
    short_name varchar(255) null,
    full_name varchar(255) null
);
CREATE TABLE d_insurance_organization(
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    code character varying(50),
    short_name character varying(255),
    full_name character varying(255)
);
CREATE TABLE d_event_type(
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name character varying(50),
    abbreviation character varying(5)
);
CREATE TABLE d_sex(
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name character varying(50),
    abbreviation character varying(5)
);
CREATE TABLE d_priority_group(
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name character varying(5)
);
CREATE TABLE d_informing_method(
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name character varying(50)
);
CREATE TABLE d_informing_type(
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name character varying(50)
);
CREATE TABLE d_health_group(
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    code character varying(50),
    name character varying(255),
    abbreviation character varying(5)
);
CREATE TABLE patients(
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    enp character varying(50) not null,
    firstname character varying(50) not null,
    secondname character varying(50) not null,
    middlename character varying(50) not null,
    birth_date INTEGER not null,
    address character varying(100),
    phone_number character varying(20),
    end_insurance_date INTEGER,
    insurance_org_id int not null,
    organization_id int not null,
    sex_id int not null,
    Foreign Key (insurance_org_id) REFERENCES d_insurance_organization (id),
    Foreign Key (organization_id) REFERENCES d_medical_organization (id),
    Foreign Key (sex_id) REFERENCES d_sex (id)
);
CREATE TABLE planning_year (
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    planned_year integer,
    planned_month integer,
    last_visit_date INTEGER,
    covid_date INTEGER,
    patient_id integer not NULL,
    planned_event_id integer,
    priority_group_id INTEGER,
    Foreign Key (patient_id) REFERENCES patients (id),
    Foreign Key (planned_event_id) REFERENCES d_event_type (id),
    Foreign Key (priority_group_id) REFERENCES d_priority_group (id)
);
CREATE TABLE approval_information( 
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    planning_year_id INTEGER not NULL,
    approve_date INTEGER NOT NULL,
    aproved_user_id INTEGER NOT NULL,
    insurance_org_id INTEGER NOT NULL,
    Foreign Key (insurance_org_id) REFERENCES d_insurance_organization (id),
    Foreign Key (planning_year_id) REFERENCES planning_year (id)
);
CREATE TABLE informing(  
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    planning_year_id INTEGER NOT NULL,
    informing_date INTEGER NOT NULL,
    informing_type_id INTEGER NOT NULL,
    informing_method_id INTEGER NOT NULL,
    Foreign Key (planning_year_id) REFERENCES planning_year (id),
    Foreign Key (informing_type_id) REFERENCES d_informing_type (id),
    Foreign Key (informing_method_id) REFERENCES d_informing_method (id)
);
CREATE TABLE progress(  
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    planning_year_id INTEGER NOT NULL,
    event_date INTEGER,
    completed_event_id INTEGER,
    health_group_id INTEGER,
    second_stage BOOLEAN NOT NULL DEFAULT false,
    executing_organization_id INTEGER,
    Foreign Key (planning_year_id) REFERENCES planning_year (id)
);
CREATE TABLE planning_group(  
    id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    planning_year_id INTEGER NOT NULL,
    accounting_serial_number INTEGER,
    insurance_org_id INTEGER NOT NULL,
    Foreign Key (insurance_org_id) REFERENCES d_insurance_organization (id),
    Foreign Key (planning_year_id) REFERENCES planning_year (id)
);
