package transfer

type PhotoCreateTransfer struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Caption string `json:"caption" form:"caption" binding:"required"`
	UserID  uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type PhotoUpdateTransfer struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photourl" form:"photourl" binding:"required"`
	UserID   uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}