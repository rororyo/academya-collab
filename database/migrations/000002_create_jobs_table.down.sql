ALTER TABLE jobseeker_skills DROP CONSTRAINT IF EXISTS jobseeker_skills_job_id_fkey;
DROP TABLE IF EXISTS jobs;
DROP TABLE IF EXISTS jobseeker_skills;