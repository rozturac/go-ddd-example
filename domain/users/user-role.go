package users

type UserRole struct {
	Name string `json:"name"`
}

var (
	UserRole_Admin = &UserRole{Name: "Admin"}
	UserRole_Guest = &UserRole{Name: "Guest"}
)
