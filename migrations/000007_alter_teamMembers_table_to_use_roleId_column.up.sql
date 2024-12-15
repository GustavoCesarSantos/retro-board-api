ALTER TABLE teamMembers ADD COLUMN roleId BIGINT;

ALTER TABLE teamMembers
DROP COLUMN role,
ALTER COLUMN roleId SET NOT NULL,
ADD CONSTRAINT fk_role FOREIGN KEY (roleId) REFERENCES teamRoles (id) ON DELETE RESTRICT;

DROP INDEX IF EXISTS idx_teamMembers_memberId_role_teamId;

CREATE INDEX idx_teamMembers_memberId_roleId_teamId
ON teamMembers (memberId, roleId, teamId);