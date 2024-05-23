CREATE TABLE "archive"(
                          "id" SERIAL PRIMARY KEY NOT NULL,
                          "doctor_availability_id" INTEGER NOT NULL,
                          "start_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                          "patient_problem" TEXT NOT NULL,
                          "end_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                          "status" VARCHAR(255) NOT NULL CHECK ("status" IN ('attended', 'cancelled', 'no_show')),
                          "payment_type" VARCHAR(255) NOT NULL CHECK ("payment_type" IN ('cash', 'card', 'insurance')),
                          "payment_amount" DOUBLE PRECISION NOT NULL,
                          "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '5 hours'),
                          "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE,
                          "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);