CREATE INDEX idx_team_members_member_id_role_team_id
ON team_members (member_id, role, team_id);