package forms

type PasswordLoginForm struct {
	// 手机号码，自定义validator
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	// 密码
	Password string `form:"password" json:"password" binding:"required,min=3,max=10"`
}
