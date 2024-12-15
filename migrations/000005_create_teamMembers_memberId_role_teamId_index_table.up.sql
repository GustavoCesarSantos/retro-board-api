CREATE INDEX idx_teamMembers_memberId_role_teamId
ON teamMembers (memberId, role, teamId);