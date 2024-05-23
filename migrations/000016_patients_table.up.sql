CREATE TABLE "patients"(
                           "id" UUID PRIMARY KEY NOT NULL,
                           "first_name" VARCHAR(50) NOT NULL,
                           "last_name" VARCHAR(50) NOT NULL,
                           "birth_date" DATE NOT NULL,
                           "gender" VARCHAR(255) NOT NULL CHECK ("gender" IN ('male', 'female', 'other')),
                           "blood_group" VARCHAR(255) NOT NULL CHECK ("blood_group" IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-')),
                           "phone_number" VARCHAR(15) NOT NULL,
                           "address" VARCHAR(250) NOT NULL,
                           "city" VARCHAR(50) NOT NULL,
                           "country" VARCHAR(50) NOT NULL,
                           "patient_problem" TEXT NOT NULL,
                           "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
                           "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE ,
                           "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);