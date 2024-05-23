CREATE TABLE "doctor_notes"(
                               "id" SERIAL PRIMARY KEY NOT NULL,
                               "appointment_id" INTEGER NOT NULL,
                               "doctor_id" UUID NOT NULL,
                               "patient_id" UUID NOT NULL,
                               "prescription" TEXT NOT NULL,
                               "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
                               "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE ,
                               "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE,
                               CONSTRAINT "doctor_notes_appointment_id_foreign" FOREIGN KEY("appointment_id") REFERENCES "archive"("id") ON DELETE CASCADE,
                               CONSTRAINT "doctor_notes_patient_id_foreign" FOREIGN KEY("patient_id") REFERENCES "patients"("id") ON DELETE CASCADE
);