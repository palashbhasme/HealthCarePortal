package services

import (
	"errors"
	"fmt"
)

var RolePermissionMap = map[string][]string{
	"doctor":       {"update_patient", "view_patient"},
	"receptionist": {"create_patient", "delete_patient", "update_patient", "view_patient"},
}

func CheckPermission(role, permission string) error {
	permissions, exists := RolePermissionMap[role]
	if !exists {
		return errors.New("invalid role")
	}

	for _, p := range permissions {
		if p == permission {
			return nil
		}
	}
	return fmt.Errorf("permission denied: %s", permission)
}
