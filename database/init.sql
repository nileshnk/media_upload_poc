CREATE USER pgadmin;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE DATABASE mediamanagement;
GRANT ALL PRIVILEGES ON DATABASE mediamanagement TO pgadmin;

-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";