ALTER TABLE team_members ADD COLUMN role VARCHAR(50);

ALTER TABLE team_members DROP CONSTRAINT fk_role;
ALTER TABLE team_members DROP COLUMN role_id;

DROP INDEX IF EXISTS idx_team_members_member_id_role_id_team_id;

CREATE INDEX idx_team_members_member_id_role_team_id
ON team_members (member_id, role, team_id);