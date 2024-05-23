CREATE TABLE IF NOT EXISTS departments
(
    id          UUID PRIMARY KEY,
    department_order   SERIAL,
    name        VARCHAR(255) NOT NULL,
    description text NOT NULL,
    image_url   text NOT NULL,
    floor_number INT NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP
    );

CREATE INDEX name_idx ON departments(lower(name));
