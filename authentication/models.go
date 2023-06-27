package authentication

type User struct {
	Id       string `json:"id" xml:"id" form:"id" query:"id"`
	Name     string `json:"name" xml:"name" form:"name" query:"name"`
	Email    string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}
