package note

import (
	"context"
	"net/http"
	"time"
	"week/internal/model"
	utils "week/internal/utlis"
	descAuth "week/pkg/auth_v1"
)

const (
	refreshTokenSecretKey  = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g0="
	accessTokenSecretKey   = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="
	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 5 * time.Minute
)

func Login(ctx context.Context, w http.ResponseWriter, req *descAuth.LoginRequest) (*descAuth.LoginResponse, error) {
	// Лезем в базу или кэш за данными пользователя
	// Сверяем хэши пароля

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: req.GetUsername(),
		// Это пример, в реальности роль должна браться из базы или кэша
		Role: "admin",
	},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return nil, err
	}

	return &descAuth.LoginResponse{RefreshToken: refreshToken}, nil
}
