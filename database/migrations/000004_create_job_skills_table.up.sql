CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE Table jobseeker_skills(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  job_id UUID NOT NULL REFERENCES jobs(id),
  skill_id UUID NOT NULL REFERENCES skills(id)
)