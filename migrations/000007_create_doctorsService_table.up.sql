CREATE TABLE IF NOT EXISTS doctor_service
(
    id UUID PRIMARY KEY NOT NULL,
    doctor_service_order SERIAL,
    doctor_id UUID REFERENCES doctors(id) ON DELETE CASCADE NOT NULL,
    specialization_id UUID REFERENCES specializations(id) ON DELETE CASCADE NOT NULL,
    online_price FLOAT NOT NULL,
    offline_price FLOAT NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
    updated_at      TIMESTAMP,
    deleted_at      TIMESTAMP
);