CREATE TABLE events (
  event_id varchar(36) PRIMARY KEY NOT NULL,
  type varchar(255) NOT NULL,
  timestamp timestamptz NOT NULL DEFAULT NOW(),
  data jsonb NOT NULL
);