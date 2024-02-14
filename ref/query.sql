CREATE TABLE role(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(9) NOT NULL,
    password VARCHAR(250) NOT NULL,
    name VARCHAR(250) NOT NULL,
    surname VARCHAR(250) NOT NULL,
    role_id SERIAL REFERENCES role NOT NULL,
    date_of_birth TIMESTAMP
);

CREATE TABLE course(
    id SERIAL PRIMARY KEY,
    name VARCHAR(250) NOT NULL,
    start_date TIMESTAMP,
    duration INTEGER NOT NULL,
    schedule VARCHAR(1000) NOT NULL,
    age_limit INTEGER,
    registration_end_date TIMESTAMP,
    address VARCHAR(250) NOT NULL,
    description VARCHAR(1000) NOT NULL,
    mentor VARCHAR(250) NOT NULL,
    format VARCHAR(200) NOT NULL,
    language VARCHAR(200) NOT NULL
);

CREATE TABLE application(
    id SERIAL PRIMARY KEY,
    user_id SERIAL REFERENCES users NOT NULL,
    course_id SERIAL REFERENCES course NOT NULL
);



