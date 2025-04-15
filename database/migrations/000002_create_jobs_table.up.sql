CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table jobs(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  title TEXT NOT NULL,
  position TEXT NOT NULL,
  description TEXT NOT NULL,
  location TEXT NOT NULL,
  salary INT,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  company_id UUID NOT NULL REFERENCES users(id)
)