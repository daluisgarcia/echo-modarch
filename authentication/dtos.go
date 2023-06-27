package authentication

type RegisterUserRequest struct {
	Name     string `json:"name" xml:"name" form:"name" query:"name"`
	Email    string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}

type UserCookieData struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
