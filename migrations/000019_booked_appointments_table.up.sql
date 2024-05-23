CREATE TABLE "booked_appointments"(
                                      "id" SERIAL PRIMARY KEY NOT NULL,
                                      "department_id" UUID NOT NULL,
                                      "doctor_id" UUID NULL,
                                      "patient_id" UUID NOT NULL,
                                      "appointment_date" DATE NOT NULL,
                                      "appointment_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                                      "duration" BIGINT NOT NULL,
                                      "key" VARCHAR(20) NOT NULL,
                                      "expires_at" TIMESTAMP(0) WITHOUT TIME ZONE,
                                      "patient_status" BOOLEAN NOT NULL DEFAULT TRUE,
                                      "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
                                      "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE ,
                                      "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);

ALTER TABLE booked_appointments
ADD CONSTRAINT unique_appointment_datetime UNIQUE (appointment_date, appointment_time);
