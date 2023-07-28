CREATE TABLE IF NOT EXISTS "users" (
  id varchar(36) NOT NULL PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  password varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "tasks" (
  id varchar(36) NOT NULL PRIMARY KEY,
  title varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  user_id varchar(36) NOT NULL,
  is_done boolean NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);