CREATE TABLE IF NOT EXISTS doctor_working_hours
(
    id SERIAL,
    doctor_id UUID REFERENCES doctors(id) ON DELETE CASCADE NOT NULL,
    day_of_week    VARCHAR(20) CHECK (day_of_week IN ('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday')) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    finish_time TIMESTAMP NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
    updated_at      TIMESTAMP,
    deleted_at      TIMESTAMP
);
