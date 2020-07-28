# Promotion service

###Description:

The repository include services to work on a promotion service.

There are two parts for the provided task.

1. Store the information of a csv file
2. Expose an API to retrieve promotions by ID

### Assumptions:
- As it's mentioned in the task, the application should expect a CSV file every 30 minutes. This is simulated by 
a ticker that clear the table and copy data from the same provided CSV file.
- The provided service is only for presentational and local usage, not production. For that purpose,
build process will be different.


### Data Model and storage:
The app uses postgresSQL as the persistence layer and use `dbr` library as an addition to
postgres driver for go.

For the migrations against database, a docker is provided to create and execute migrations.

Data model:
```sql
    row_id  INTEGER,
    id      VARCHAR(50),
    price   DECIMAL,
    expiration_date timestamp
``` 
row_id is added to let the API retrieve promotions by row number they have in CSV file

## Services:
The app provides 3 services
- API
- CSV service which only imports the data from CSV to storage
- All which include both services (API and CSV)

Suggestion: To have a better performance of copying data from CSV to storage, if postgres
is used, it is better to use copy command of postgres and a tool (provided by docker)
to copy the data.

### Prerequisite
- Docker and Docker-compose
- make command (unix linux environment)

### How to run
As it is recommended by 12-Factor app, services use environment variables for configs.
required env vars are provided in .env file

To run the services there are 3 make commands provided
- run-all-services - this command runs ALL 
- run-api-service - This command runs only the web API service
- run-csv-service - This command runs only the process for copying csv file to storage


