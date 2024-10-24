package service

import(
	"auto_slip_api/model"
	"auto_slip_api/pkg/utils"
	"auto_slip_api/pkg/custom"
)

type LoginService struct{}

// ログイン
func (s *LoginService) Login(loginUser model.User) (string, error) {

	// userを取得
	user,err := model.GetUserById(loginUser.UserId)
	if err != nil {
		return "",custom.NewErr(custom.ErrTypeNoResourceExist)	// ユーザーが存在しないよエラー
	}
	if user.UserId == "" {
		return "",custom.NewErr(custom.ErrTypeNoResourceExist)	// ユーザーが存在しないよエラー
	}
	// パスワードの比較を行う
	login := model.CheckPasswordHash(loginUser.Password, user.Password)
	var token string
	if login {
		token,err = utils.GenerateToken(user.UserId)
		if err != nil {
			return "",custom.NewErr(custom.ErrTypeGenTokenFailed)	// トークン生成失敗エラー
		}
	}else{
		return "",custom.NewErr(custom.ErrTypePermissionDenied)
	}
	return token,nil
}

func(s *LoginService) CreateUser() error {
	
	// ユーザー情報を登録
	err := model.RegisterUser()
	if err != nil {
		return err	// ユーザー登録失敗エラー
	}
	return nil
}

