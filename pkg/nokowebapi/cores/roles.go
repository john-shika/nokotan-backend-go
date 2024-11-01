package cores

type RoleTyped string

const (
	RoleAdmin RoleTyped = "admin"
	RoleUser  RoleTyped = "user"
	RoleGuest RoleTyped = "guest"
)
