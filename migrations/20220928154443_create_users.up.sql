/* migrate create -ext sql -dir migrations create_users */
/* migrate  -path migrations -database "postgres://postgres:Cain-666@localhost/restapi_dev?sslmode=disable" up */
CREATE TABLE users (
  id serial not null primary key,
  email varchar not null unique,
  encrypted_password varchar not null
);