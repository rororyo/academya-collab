CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table IF NOT EXISTS skills(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  skill_name TEXT NOT NULL
)