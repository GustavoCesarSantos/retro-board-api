ALTER TABLE teamMembers ADD COLUMN role VARCHAR(50);

ALTER TABLE teamMembers DROP CONSTRAINT fk_role;
ALTER TABLE teamMembers DROP COLUMN roleId;

DROP INDEX IF EXISTS idx_teamMembers_memberId_roleId_teamId;

CREATE INDEX idx_teamMembers_memberId_role_teamId
ON teamMembers (memberId, role, teamId);