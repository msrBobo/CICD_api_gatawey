CREATE TYPE role_type AS ENUM('superadmin', 'admin');

-- /*admins table*/
CREATE TABLE IF NOT EXISTS admins (
    id UUID NOT NULL PRIMARY KEY,
    admin_order SERIAL,
    role role_type NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    birth_date DATE NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(50) NOT NULL,
    gender gender_type NOT NULL,
    salary FLOAT NOT NULL DEFAULT 0.0, 
    biography TEXT NOT NULL,
    start_work_year DATE NOT NULL,
    end_work_year DATE,
    work_years INTEGER DEFAULT 0,
    refresh_token TEXT NOT NULL,
    image_url VARCHAR(200),
    created_at TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX admin_unique_phone_number_deleted_at_null_idx ON admins(phone_number) WHERE deleted_at IS NULL; --unique admin phone number, users with non-soft-deleted accounts cannot register again with the same phone number.
CREATE INDEX deleted_at_idx2 ON admins(deleted_at);

---------------------------------------------------------------------------------------------------------------------------

-- For admins table
INSERT INTO admins (id, admin_order, role, first_name, last_name, birth_date, phone_number, email, password, gender, salary, biography, start_work_year, refresh_token, created_at, image_url)
VALUES
  ('123e4567-e89b-12d3-a456-426614174001', 1, 'admin',    'first_name1', 'last_name1', '2000-01-01', 'phone_number1', 'email1@example.com', 'password1', 'male', 1000.00, '1990-05-15', '2000-01-01', 'refresh_token1', CURRENT_TIMESTAMP, 'https://example.com/admin1.jpg'),
  ('123e4567-e89b-12d3-a456-426614174002', 2, 'superadmin','first_name2', 'last_name2', '2000-01-02', 'phone_number2', 'email2@example.com', 'password2', 'male', 2000.00, '1990-05-15', '2000-01-02', 'refresh_token2', CURRENT_TIMESTAMP, 'https://example.com/superadmin2.jpg'),
  ('123e4567-e89b-12d3-a456-426614174003', 3, 'admin', 'first_name3', 'last_name3', '2000-01-03', 'phone_number3', 'email3@example.com', 'password3', 'male', 3000.00, '1990-05-15', '2000-01-03', 'refresh_token3', CURRENT_TIMESTAMP, 'https://example.com/admin3.jpg'),
  ('123e4567-e89b-12d3-a456-426614174004', 4, 'admin', 'first_name4', 'last_name4', '2000-01-04', 'phone_number4', 'email4@example.com', 'password4', 'male', 4000.00, '1990-05-15', '2000-01-04', 'refresh_token4', CURRENT_TIMESTAMP, 'https://example.com/admin4.jpg'),
  ('123e4567-e89b-12d3-a456-426614174005', 5, 'superadmin', 'first_name5', 'last_name5', '2000-01-05', 'phone_number5', 'email5@example.com', 'password5', 'male', 5000.00, '1990-05-15', '2000-01-05', 'refresh_token5', CURRENT_TIMESTAMP, 'https://example.com/superadmin5.jpg'),
  ('123e4567-e89b-12d3-a456-426614174006', 6, 'superadmin', 'first_name6', 'last_name6', '2000-01-06', 'phone_number6', 'email6@example.com', 'password6', 'male', 6000.00, '1990-05-15', '2000-01-06', 'refresh_token6', CURRENT_TIMESTAMP, 'https://example.com/superadmin6.jpg'),
  ('123e4567-e89b-12d3-a456-426614174007', 7, 'admin', 'first_name7', 'last_name7', '2000-01-07', 'phone_number7', 'email7@example.com', 'password7', 'male', 7000.00, '1990-05-15', '2000-01-07', 'refresh_token7', CURRENT_TIMESTAMP, 'https://example.com/admin7.jpg'),
  ('123e4567-e89b-12d3-a456-426614174008', 8, 'superadmin', 'first_name8', 'last_name8', '2000-01-08', 'phone_number8', 'email8@example.com', 'password8', 'male', 8000.00, '1990-05-15', '2000-01-08', 'refresh_token8', CURRENT_TIMESTAMP, 'https://example.com/superadmin8.jpg'),
  ('123e4567-e89b-12d3-a456-426614174009', 9, 'admin', 'first_name9', 'last_name9', '2000-01-09', 'phone_number9', 'email9@example.com', 'password9', 'male', 9000.00, '1990-05-15', '2000-01-09', 'refresh_token9', CURRENT_TIMESTAMP, 'https://example.com/admin9.jpg'),
  ('123e4567-e89b-12d3-a456-426614174010', 10, 'superadmin', 'first_name10', 'last_name10', '2000-01-10', 'phone_number10', 'email10@example.com', 'password10', 'male', 10000.00, '1990-05-15', '2000-01-10', 'refresh_token10', CURRENT_TIMESTAMP, 'https://example.com/superadmin10.jpg');