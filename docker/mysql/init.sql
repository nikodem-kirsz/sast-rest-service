-- Create the database
CREATE DATABASE sast_database;
USE sast_database;

-- Create the table
CREATE TABLE mysql_report_models (
    id INT AUTO_INCREMENT PRIMARY KEY,
    uuid VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    time DATETIME,
    report_content TEXT
);