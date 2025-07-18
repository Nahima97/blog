package config

func RolePermission(userRole string, action string) bool {

	if userRole == "admin" {
		return true
	} else if userRole == "user" {
		return action == "create:post" || action == "read:post" || action == "update:post" || action == "delete:post"
	} else {
		return false
	}

}
