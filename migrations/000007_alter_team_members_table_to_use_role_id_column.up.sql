ALTER TABLE team_members ADD COLUMN role_id BIGINT;

ALTER TABLE team_members
DROP COLUMN role,
ALTER COLUMN role_id SET NOT NULL,
ADD CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES team_roles (id) ON DELETE RESTRICT;

DROP INDEX IF EXISTS idx_team_members_member_id_role_team_id;

CREATE INDEX idx_team_members_member_id_role_id_team_id
ON team_members (member_id, role_id, team_id);