package user_management_model

import (
    "errors"
    "gorm.io/gorm"
    "time"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

package user_management_model

import (
    "errors"
    "gorm.io/gorm"
    "time"
)

// CreateSuperAdmin creates a superadmin user, assigns a role, and grants all permissions
func CreateSuperAdmin(db *gorm.DB, superAdmin Admin, permissions []Permission) error {
    return db.Transaction(func(tx *gorm.DB) error {
        // Check if superadmin already exists
        var existingAdmin Admin
        if err := tx.Where("email = ?", superAdmin.Email).First(&existingAdmin).Error; err == nil {
            return errors.New("superadmin with this email already exists")
        }

        // Create superadmin user
        superAdmin.CreatedAt = time.Now().Unix()
        superAdmin.UpdatedAt = time.Now().Unix()
        if err := tx.Create(&superAdmin).Error; err != nil {
            return err
        }

        // Create a role for superadmin
        superAdminRole := Role{
            Name:      "SuperAdmin",
            CreatedAt: time.Now().Unix(),
            UpdatedAt: time.Now().Unix(),
        }

        if err := tx.Create(&superAdminRole).Error; err != nil {
            return err
        }

        // Associate all permissions to the role
        if err := tx.Model(&superAdminRole).Association("Permissions").Append(permissions); err != nil {
            return err
        }

        // Associate the role with the superadmin
        if err := tx.Model(&superAdmin).Association("Roleable").Append(superAdminRole); err != nil {
            return err
        }

        return nil
    })
}



