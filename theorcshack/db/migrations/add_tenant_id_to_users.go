package migrations

import (
	"gorm.io/gorm"
)

func AddTenantIDToUsers(db *gorm.DB) {
	db.Exec("ALTER TABLE users ADD COLUMN tenant_id uuid;")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_users_tenant_id ON users (tenant_id);")
}
