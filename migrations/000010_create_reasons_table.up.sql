CREATE TABLE IF NOT EXISTS reasons
(
    id UUID NOT NULL,
    name VARCHAR(200) NOT NULL,
    specialization_id UUID REFERENCES specializations(id) ON DELETE CASCADE NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
    updated_at    TIMESTAMP ,
    deleted_at    TIMESTAMP
);