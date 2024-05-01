package requestmodels_auth_apigw

type UserSignUpReq struct {
	Name            string `json:"name" validate:"required,gte=3,lte=30"`
	UserName        string `json:"username" validate:"required,gte=3,lte=30"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,gte=3,lte=30"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type OtpVerification struct {
	Otp string `json:"otp" validate:"required,len=4,number"`
}

type UserLoginReq struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=30"`
}

type ForgotPasswordReq struct {
	Email string `json:"email"    validate:"required,email"`
}

type ForgotPasswordData struct {
	Otp             string `json:"otp" validate:"required,len=4,number"`
	Password        string `json:"password" validate:"required,gte=3,lte=30"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type EditUserProfile struct {
	Name     string `json:"name" validate:"required,gte=3,lte=25"`
	UserName string `json:"username" validate:"required,gte=3,lte=30"`
	Bio      string `json:"bio" validate:"lte=50"`
	Links    string `json:"links" validate:"lte=25"`

	UserId string
}
