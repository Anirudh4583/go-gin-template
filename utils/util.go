package util

import "github.com/Anirudh4583/go-gin-template/pkg/setting"

func Setup() {
	jwtSecret = []byte(setting.Config.AppJwtSecret)
}
