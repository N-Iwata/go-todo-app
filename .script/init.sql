create table users(
  id SERIAL PRIMARY KEY,
  uuid varchar(64) NOT NULL UNIQUE,
  name VARCHAR(255),
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE sessions(
  id SERIAL NOT NULL,
  uuid VARCHAR(64) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE,
  user_id INTEGER REFERENCES users(id),
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE todos(
  id SERIAL PRIMARY KEY,
  content TEXT,
  user_id INTEGER REFERENCES users(id),
  created_at TIMESTAMP NOT NULL
);