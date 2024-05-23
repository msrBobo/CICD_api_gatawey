CREATE TABLE IF NOT EXISTS doctors
(
    id              UUID PRIMARY KEY,
    doctor_order SERIAL,
    first_name      varchar(200) NOT NULL,
    last_name       varchar(200) NOT NULL,
    image_url       varchar(200) NOT NULL,
    gender          VARCHAR(6) CHECK (gender IN ('male', 'female')) NOT NULL,
    birth_date DATE NOT NULL ,
    phone_number    VARCHAR(15) NOT NULL,
    email           VARCHAR(100) NOT NULL,
    password        text NOT NULL,
    address         varchar(255) NOT NULL,
    city TEXT NOT NULL,
    country TEXT NOT NULL,
    salary          FLOAT NOT NULL,
    biography       TEXT NOT NULL,
    start_work_date DATE NOT NULL,
    end_work_date   DATE,
    work_years      INT NOT NULL,
    department_id   UUID REFERENCES departments(id) ON DELETE CASCADE NOT NULL,
    room_number INT NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
    updated_at      TIMESTAMP,
    deleted_at      TIMESTAMP
    );

CREATE UNIQUE INDEX unique_email_deleted_at_null_idx ON doctors(email) WHERE deleted_at IS NULL ;

CREATE INDEX department_id_idx ON doctors(department_id);
CREATE INDEX first_name_idx ON doctors(first_name);
CREATE INDEX last_name_idx ON doctors(last_name);
CREATE INDEX work_years_idx ON doctors(work_years);