package user

// UserMapper map
type UserMapper struct {
	ID       uint   `json:"id,string,omitempty"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
	Telepon  string `json:"telepon"`
}

type LoginResponseMapper struct {
	ID       uint   `json:"id,string"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func toTable(u UserMapper) User {
	return User{Nama: u.Nama, Username: u.Username, Password: u.Password, Telepon: u.Telepon}
}

func toResponse(u User) UserMapper {
	return UserMapper{ID: u.ID, Nama: u.Nama, Username: u.Username, Password: u.Password, Telepon: u.Telepon}
}

func toMultipleResponse(u []User) []UserMapper {
	ret := make([]UserMapper, len(u))

	for k, v := range u {
		ret[k] = toResponse(v)
	}

	return ret
}

func toLoginResponse(u User) LoginResponseMapper {
	return LoginResponseMapper{ID: u.ID, Nama: u.Nama, Username: u.Username, Token: "Put token here"}
}
