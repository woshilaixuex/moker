// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`
	Username string `json:"username"`
}

type VerCodeResp struct {
	Base64 string `json:"b64s"`
	Key    string `json:"key"`
}

type LoginReq struct {
	Account  string `json:"account"`
	PassWord string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RegisterReq struct {
	VerCode  string `json:"ver_code"`
	VerKey   string `json:"ver_key"`
	Account  string `json:"account"`
	PassWord string `json:"password"`
	Role     string `json:"role"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
	Role         string `json:"role"`
}

type Role struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	Name    string `json:"name"`
	Major   string `json:"major"`
	Faculty string `json:"faculty"`
	School  string `json:"school"`
}

type GetUserInfoReq struct {
}

type GetUserInfoResp struct {
	Username string `json:"username"`
	Role     Role   `json:"role"`
}

type UpdateStudentInfoReq struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Major   string `json:"major"`
	Faculty string `json:"faculty"`
	School  string `json:"school"`
}

type UpdateStudentInfoResp struct {
	Role Role `json:"role"`
}

type UpdateTeacherInfoReq struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Faculty string `json:"faculty"`
	School  string `json:"school"`
}

type UpdateTeacherInfoResp struct {
	Role Role `json:"role"`
}

type DeleteUserInfoReq struct {
	StuId int64 `json:"stu_id"`
}

type DeleteUserInfoResp struct {
	IsOk bool `json:"isOk"`
}
