-- Version: 1.01
-- Description: Create table users
CREATE TABLE users (
    user_id UUID NOT NULL,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    roles TEXT [] NOT NULL,
    password_hash TEXT NOT NULL,
    department TEXT NULL,
    enabled BOOLEAN NOT NULL,
    date_created TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id)
);

-- Version: 1.04
-- Description: Create table scrums
CREATE TABLE scrums (
    scrum_id UUID NOT NULL,
    name TEXT NOT NULL,
    time INTEGER NOT NULL,
    color TEXT NOT NULL,
    attendees TEXT NULL,
    user_id UUID NOT NULL,
    date_created TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL,
    PRIMARY KEY (scrum_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);