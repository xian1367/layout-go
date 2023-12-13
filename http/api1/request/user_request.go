package request

type UserStoreRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func UserStore(data interface{}) (errs map[string]string) {
	return errs
}

type UserUpdateRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func UserUpdate(data interface{}) (errs map[string]string) {
	return errs
}
