# Use an official MySQL image from Docker Hub
FROM mysql:latest

# Set the root password for the MySQL server
ENV MYSQL_ROOT_PASSWORD password

# Copy a SQL script to initialize the database
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the default MySQL port
EXPOSE 3306
