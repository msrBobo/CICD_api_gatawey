INSERT INTO departments (id, department_order, name, description, image_url, floor_number, created_at, updated_at, deleted_at, short_description) VALUES
('123e4567-e89b-12d3-a456-426614174001', 1, 'Human Resources', 'Responsible for managing employee relations, recruitment, and training.', 'https://example.com/hr_image.jpg', 1, '2024-05-09 08:00:00', NULL, NULL, 'HR'),
('123e4567-e89b-12d3-a456-426614174002', 2, 'Finance', 'Responsible for managing financial activities, budgeting, and accounting.', 'https://example.com/finance_image.jpg', 2, '2024-05-09 09:00:00', NULL, NULL, 'Finance'),
('123e4567-e89b-12d3-a456-426614174003', 3, 'Marketing', 'Responsible for creating and implementing marketing strategies.', 'https://example.com/marketing_image.jpg', 1, '2024-05-09 10:00:00', NULL, NULL, 'Marketing'),
('123e4567-e89b-12d3-a456-426614174004', 4, 'IT', 'Responsible for managing information technology resources and systems.', 'https://example.com/it_image.jpg', 3, '2024-05-09 11:00:00', NULL, NULL, 'IT'),
('123e4567-e89b-12d3-a456-426614174005', 5, 'Sales', 'Responsible for selling products or services to customers.', 'https://example.com/sales_image.jpg', 1, '2024-05-09 12:00:00', NULL, NULL, 'Sales'),
('123e4567-e89b-12d3-a456-426614174006', 6, 'Customer Service', 'Responsible for assisting customers with inquiries and complaints.', 'https://example.com/customer_service_image.jpg', 2, '2024-05-09 13:00:00', NULL, NULL, 'Customer Service'),
('123e4567-e89b-12d3-a456-426614174007', 7, 'Operations', 'Responsible for overseeing daily business operations and processes.', 'https://example.com/operations_image.jpg', 3, '2024-05-09 14:00:00', NULL, NULL, 'Operations'),
('123e4567-e89b-12d3-a456-426614174008', 8, 'Research and Development', 'Responsible for researching and developing new products or services.', 'https://example.com/rnd_image.jpg', 2, '2024-05-09 15:00:00', NULL, NULL, 'R&D'),
('123e4567-e89b-12d3-a456-426614174009', 9, 'Legal', 'Responsible for providing legal advice and ensuring compliance with laws and regulations.', 'https://example.com/legal_image.jpg', 1, '2024-05-09 16:00:00', NULL, NULL, 'Legal'),
('123e4567-e89b-12d3-a456-426614174010', 10, 'Public Relations', 'Responsible for managing the public image and perception of the company.', 'https://example.com/pr_image.jpg', 3, '2024-05-09 17:00:00', NULL, NULL, 'PR');

----------------------------------------------------------------------------------------------------------------------------------------------

INSERT INTO doctors (id, doctor_order, first_name, last_name, image_url, gender, birth_date, phone_number, email, password, address, city, country, salary, biography, start_work_date, end_work_date, work_years, department_id, room_number, created_at, updated_at, deleted_at) VALUES
('123e4567-e89b-12d3-a456-426614274001', 1, 'John', 'Doe', 'imageUrl', 'male', '1985-04-15', '+1234567890', 'john.doe@example.com', 'hashed_password_1', '123 Main St', 'New York', 'USA', 100000.00, 'Dr. John Doe has been practicing medicine for over a decade, specializing in cardiology.', '2010-05-01', NULL, 15, '123e4567-e89b-12d3-a456-426614174001', 101, '2024-05-09 08:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274002', 2, 'Jane', 'Smith', 'imageUrl', 'female', '1980-08-25', '+1987654321', 'jane.smith@example.com', 'hashed_password_2', '456 Elm St', 'Los Angeles', 'USA', 95000.00, 'Dr. Jane Smith specializes in pediatric care and has received numerous awards for her contributions to child healthcare.', '2005-09-15', NULL, 16, '123e4567-e89b-12d3-a456-426614174002', 102, '2024-05-09 09:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274003', 3, 'Michael', 'Johnson', 'imageUrl', 'male', '1976-11-10', '+1122334455', 'michael.johnson@example.com', 'hashed_password_3', '789 Oak St', 'Chicago', 'USA', 110000.00, 'Dr. Michael Johnson is a renowned surgeon specializing in orthopedics, known for his innovative surgical techniques.', '2002-03-20', NULL, 22, '123e4567-e89b-12d3-a456-426614174003', 103, '2024-05-09 10:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274004', 4, 'Emily', 'Brown', 'imageUrl', 'female', '1990-02-28', '+1654321876', 'emily.brown@example.com', 'hashed_password_4', '321 Pine St', 'Houston', 'USA', 90000.00, 'Dr. Emily Brown is an experienced family physician dedicated to providing comprehensive healthcare to patients of all ages.', '2015-07-10', NULL, 9, '123e4567-e89b-12d3-a456-426614174004', 104, '2024-05-09 11:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274005', 5, 'David', 'Wilson', 'imageUrl', 'male', '1988-09-05', '+1789456123', 'david.wilson@example.com', 'hashed_password_5', '555 Maple St', 'Dallas', 'USA', 105000.00, 'Dr. David Wilson specializes in neurology and has conducted groundbreaking research on brain disorders.', '2013-02-15', NULL, 11, '123e4567-e89b-12d3-a456-426614174005', 105, '2024-05-09 12:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274006', 6, 'Sarah', 'Miller', 'imageUrl', 'female', '1983-06-20', '+1876543210', 'sarah.miller@example.com', 'hashed_password_6', '888 Oak St', 'San Francisco', 'USA', 98000.00, 'Dr. Sarah Miller is an expert in dermatology, offering advanced treatments for skin conditions.', '2008-11-30', NULL, 16, '123e4567-e89b-12d3-a456-426614174006', 106, '2024-05-09 13:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274007', 7, 'Christopher', 'Thompson', 'imageUrl', 'male', '1982-03-12', '+1432156789', 'christopher.thompson@example.com', 'hashed_password_7', '999 Pine St', 'Seattle', 'USA', 97000.00, 'Dr. Christopher Thompson is an accomplished oncologist known for his compassionate patient care and research in cancer treatment.', '2007-06-25', NULL, 14, '123e4567-e89b-12d3-a456-426614174007', 107, '2024-05-09 14:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274008', 8, 'Amanda', 'Young', 'imageUrl', 'female', '1975-12-30', '+1567890123', 'amanda.young@example.com', 'hashed_password_8', '111 Elm St', 'Boston', 'USA', 120000.00, 'Dr. Amanda Young is a leading psychiatrist specializing in mood disorders and addiction treatment.', '2000-08-05', NULL, 24, '123e4567-e89b-12d3-a456-426614174008', 108, '2024-05-09 15:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274009', 9, 'Kevin', 'Clark', 'imageUrl', 'male', '1970-07-18', '+1908765432', 'kevin.clark@example.com', 'hashed_password_9', '222 Oak St', 'Miami', 'USA', 115000.00, 'Dr. Kevin Clark is a respected cardiologist with extensive experience in treating heart diseases.', '1995-12-10', NULL, 29, '123e4567-e89b-12d3-a456-426614174009', 109, '2024-05-09 16:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274010', 10, 'Rachel', 'Adams', 'imageUrl', 'female', '1987-01-08', '+1209876543', 'rachel.adams@example.com', 'hashed_password_10', '333 Maple St', 'Atlanta', 'USA', 95000.00, 'Dr. Rachel Adams is a skilled obstetrician-gynecologist dedicated to women’s health and wellness.', '2011-04-20', NULL, 13, '123e4567-e89b-12d3-a456-426614174010', 110, '2024-05-09 17:00:00', NULL, NULL);

----------------------------------------------------------------------------------------------------------------------------------------------

INSERT INTO specializations (id, specializations_order, name, description, department_id, created_at, updated_at, deleted_at, image_url) VALUES
('123e4567-e89b-12d3-a456-426614375001', 1, 'Cardiology', 'Specializing in the diagnosis, treatment, and prevention of heart diseases and disorders.', '123e4567-e89b-12d3-a456-426614174001', '2024-05-09 08:00:00', NULL, NULL, 'https://example.com/cardiology_image.jpg'),
('123e4567-e89b-12d3-a456-426614375002', 2, 'Pediatrics', 'Specializing in medical care for infants, children, and adolescents.', '123e4567-e89b-12d3-a456-426614174003', '2024-05-09 09:00:00', NULL, NULL, 'https://example.com/pediatrics_image.jpg'),
('123e4567-e89b-12d3-a456-426614375003', 3, 'Orthopedics', 'Specializing in the diagnosis and treatment of disorders related to the musculoskeletal system.', '123e4567-e89b-12d3-a456-426614174003', '2024-05-09 10:00:00', NULL, NULL, 'https://example.com/orthopedics_image.jpg'),
('123e4567-e89b-12d3-a456-426614375004', 4, 'Dermatology', 'Specializing in the diagnosis and treatment of skin, hair, and nail disorders.', '123e4567-e89b-12d3-a456-426614174006', '2024-05-09 11:00:00', NULL, NULL, 'https://example.com/dermatology_image.jpg'),
('123e4567-e89b-12d3-a456-426614375005', 5, 'Neurology', 'Specializing in the diagnosis and treatment of disorders of the nervous system.', '123e4567-e89b-12d3-a456-426614174005', '2024-05-09 12:00:00', NULL, NULL, 'https://example.com/neurology_image.jpg'),
('123e4567-e89b-12d3-a456-426614375006', 6, 'Oncology', 'Specializing in the diagnosis and treatment of cancer.', '123e4567-e89b-12d3-a456-426614174007', '2024-05-09 13:00:00', NULL, NULL, 'https://example.com/oncology_image.jpg'),
('123e4567-e89b-12d3-a456-426614375007', 7, 'Psychiatry', 'Specializing in the diagnosis, treatment, and prevention of mental disorders.', '123e4567-e89b-12d3-a456-426614174008', '2024-05-09 14:00:00', NULL, NULL, 'https://example.com/psychiatry_image.jpg'),
('123e4567-e89b-12d3-a456-426614375008', 8, 'Cardiothoracic Surgery', 'Specializing in surgical procedures of the heart and chest.', '123e4567-e89b-12d3-a456-426614174001', '2024-05-09 15:00:00', NULL, NULL, 'https://example.com/cardiothoracic_surgery_image.jpg'),
('123e4567-e89b-12d3-a456-426614375009', 9, 'Family Medicine', 'Specializing in primary care for individuals and families of all ages.', '123e4567-e89b-12d3-a456-426614174004', '2024-05-09 16:00:00', NULL, NULL, 'https://example.com/family_medicine_image.jpg'),
('123e4567-e89b-12d3-a456-426614375010', 10, 'Obstetrics and Gynecology', 'Specializing in womens reproductive health and childbirth.', '123e4567-e89b-12d3-a456-426614174002', '2024-05-09 17:00:00', NULL, NULL, 'https://example.com/obstetrics_gynecology_image.jpg');

----------------------------------------------------------------------------------------------------------------------------------------------

INSERT INTO doctor_working_hours (doctor_id, day_of_week, start_time, finish_time, created_at, updated_at, deleted_at) VALUES
('123e4567-e89b-12d3-a456-426614274001', 'Monday', '2024-05-09 08:00:00', '2024-05-09 17:00:00', '2024-05-09 08:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274002', 'Tuesday', '2024-05-09 08:00:00', '2024-05-09 17:00:00', '2024-05-09 08:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274003', 'Wednesday', '2024-05-09 08:00:00', '2024-05-09 17:00:00', '2024-05-09 08:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274004', 'Thursday', '2024-05-09 08:00:00', '2024-05-09 17:00:00', '2024-05-09 08:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274005', 'Friday', '2024-05-09 08:00:00', '2024-05-09 17:00:00', '2024-05-09 08:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274006', 'Saturday', '2024-05-09 09:00:00', '2024-05-09 12:00:00', '2024-05-09 08:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274007', 'Monday', '2024-05-09 09:00:00', '2024-05-09 18:00:00', '2024-05-09 09:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274008', 'Tuesday', '2024-05-09 09:00:00', '2024-05-09 18:00:00', '2024-05-09 09:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274009', 'Wednesday', '2024-05-09 09:00:00', '2024-05-09 18:00:00', '2024-05-09 09:00:00', NULL, NULL),
('123e4567-e89b-12d3-a456-426614274010', 'Thursday', '2024-05-09 09:00:00', '2024-05-09 18:00:00', '2024-05-09 09:00:00', NULL, NULL);

-- ----------------------------------------------------------------------------------------------------------------------------------------------

INSERT INTO doctor_service (id, doctor_id, specialization_id, online_price, offline_price, created_at, updated_at, deleted_at, name, duration) VALUES
('123e4567-e89b-12d3-a456-436614274001', '123e4567-e89b-12d3-a456-426614274001', '123e4567-e89b-12d3-a456-426614375001', 100.00, 150.00, '2024-05-09 08:00:00', NULL, NULL, 'General Consultation', '2024-05-09 00:30:00'),
('123e4567-e89b-12d3-a456-436614274002', '123e4567-e89b-12d3-a456-426614274002', '123e4567-e89b-12d3-a456-426614375002', 120.00, 180.00, '2024-05-09 09:00:00', NULL, NULL, 'Cardiac Examination', '2024-05-09 00:40:00'),
('123e4567-e89b-12d3-a456-436614274003', '123e4567-e89b-12d3-a456-426614274003', '123e4567-e89b-12d3-a456-426614375003', 90.00, 140.00, '2024-05-09 10:00:00', NULL, NULL, 'Dental Cleaning', '2024-05-09 00:45:00'),
('123e4567-e89b-12d3-a456-436614274004', '123e4567-e89b-12d3-a456-426614274004', '123e4567-e89b-12d3-a456-426614375004', 110.00, 160.00, '2024-05-09 11:00:00', NULL, NULL, 'Eye Examination', '2024-05-09 00:30:00'),
('123e4567-e89b-12d3-a456-436614274005', '123e4567-e89b-12d3-a456-426614274005', '123e4567-e89b-12d3-a456-426614375005', 95.00, 145.00, '2024-05-09 12:00:00', NULL, NULL, 'Physical Therapy', '2024-05-09 01:15:00'),
('123e4567-e89b-12d3-a456-436614274006', '123e4567-e89b-12d3-a456-426614274006', '123e4567-e89b-12d3-a456-426614375006', 105.00, 155.00, '2024-05-09 13:00:00', NULL, NULL, 'Psychological Counseling', '2024-05-09 01:00:00'),
('123e4567-e89b-12d3-a456-436614274007', '123e4567-e89b-12d3-a456-426614274007', '123e4567-e89b-12d3-a456-426614375007', 115.00, 165.00, '2024-05-09 14:00:00', NULL, NULL, 'Surgical Procedure', '2024-05-09 02:00:00'),
('123e4567-e89b-12d3-a456-436614274008', '123e4567-e89b-12d3-a456-426614274008', '123e4567-e89b-12d3-a456-426614375008', 125.00, 175.00, '2024-05-09 15:00:00', NULL, NULL, 'Ultrasound Scan', '2024-05-09 00:45:00'),
('123e4567-e89b-12d3-a456-436614274009', '123e4567-e89b-12d3-a456-426614274009', '123e4567-e89b-12d3-a456-426614375009', 85.00, 135.00, '2024-05-09 16:00:00', NULL, NULL, 'X-ray Examination', '2024-05-09 00:30:00'),
('123e4567-e89b-12d3-a456-436614274010', '123e4567-e89b-12d3-a456-426614274010', '123e4567-e89b-12d3-a456-426614375010', 130.00, 180.00, '2024-05-09 17:00:00', NULL, NULL, 'MRI Scan', '2024-05-09 01:30:00');


-- ----------------------------------------------------------------------------------------------------------------------------------------------

INSERT INTO reasons (id, name, specialization_id, created_at, updated_at, deleted_at, image_url) VALUES
('123e4567-e89b-12d3-a456-426614184001', 'Headache', '123e4567-e89b-12d3-a456-426614375001', CURRENT_TIMESTAMP, NULL, NULL, 'https://example.com/headache_image.jpg'),
('123e4567-e89b-12d3-a456-426614184002', 'Fever', '123e4567-e89b-12d3-a456-426614375002', CURRENT_TIMESTAMP, NULL, NULL, 'https://example.com/fever_image.jpg'),
('123e4567-e89b-12d3-a456-426614184003', 'Back Pain', '123e4567-e89b-12d3-a456-426614375003', CURRENT_TIMESTAMP, NULL, NULL, 'https://example.com/back_pain_image.jpg'),
('123e4567-e89b-12d3-a456-426614184004', 'Cough', '123e4567-e89b-12d3-a456-426614375004', CURRENT_TIMESTAMP, NULL, NULL, 'https://example.com/cough_image.jpg'),
('123e4567-e89b-12d3-a456-426614184005', 'Sore Throat', '123e4567-e89b-12d3-a456-426614375005', CURRENT_TIMESTAMP, NULL, NULL, 'https://example.com/sore_throat_image.jpg'),
('123e4567-e89b-12d3-a456-426614184006', 'Fatigue', '123e4567-e89b-12d3-a456-426614375006', CURRENT_TIMESTAMP, NULL, NULL, 'https://example.com/fatigue_image.jpg'),
('123e4567-e89b-12d3-a456-426614184007', 'Stomachache', '123e4567-e89b-12d3-a456-426614375007', CURRENT_TIMESTAMP, NULL, NULL, 'https://example.com/stomachache_image.jpg'),
('123e4567-e89b-12d3-a456-426614184008', 'Allergy', '123e4567-e89b-12d3-a456-426614375008', CURRENT_TIMESTAMP, NULL, NULL, 'https://example.com/allergy_image.jpg'),
('123e4567-e89b-12d3-a456-426614184009', 'Anxiety', '123e4567-e89b-12d3-a456-426614375009', CURRENT_TIMESTAMP, NULL, NULL, 'https://example.com/anxiety_image.jpg'),
('123e4567-e89b-12d3-a456-426614184010', 'Depression', '123e4567-e89b-12d3-a456-426614375010', CURRENT_TIMESTAMP, NULL, NULL, 'https://example.com/depression_image.jpg');
