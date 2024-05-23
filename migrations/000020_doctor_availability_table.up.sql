CREATE TABLE "doctor_availability"(
                                      "id" SERIAL PRIMARY KEY NOT NULL,
                                      "department_id" UUID NOT NULL,
                                      "doctor_id" UUID NOT NULL,
                                      "doctor_date" DATE NOT NULL,
                                      "start_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                                      "end_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                                      "status" VARCHAR(255) NOT NULL CHECK ("status" IN ('available', 'unavailable')),
                                      "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
                                      "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE ,
                                      "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);

ALTER TABLE "archive" ADD CONSTRAINT "archive_doctor_availability_id_foreign" FOREIGN KEY("doctor_availability_id") REFERENCES "doctor_availability"("id");