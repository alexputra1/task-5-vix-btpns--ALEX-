package transfer

type UserCreateTransfer struct {
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required,min=6"`
}

type UserUpdateTransfer struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required,min:6"`
}