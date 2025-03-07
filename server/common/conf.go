package common

import (
	"strings"

	"github.com/ArtalkJS/ArtalkGo/internal/config"
	"github.com/ArtalkJS/ArtalkGo/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func GetApiVersionDataMap() Map {
	return Map{
		"app":            "artalk-go",
		"version":        strings.TrimPrefix(config.Version, "v"),
		"commit_hash":    config.CommitHash,
		"fe_min_version": strings.TrimPrefix(config.FeMinVersion, "v"),
	}
}

func GetApiPublicConfDataMap(c *fiber.Ctx) Map {
	isAdmin := CheckIsAdminReq(c)
	imgUpload := config.Instance.ImgUpload.Enabled
	if isAdmin {
		imgUpload = true // 管理员始终允许上传图片
	}

	frontendConfSrc := config.Instance.Frontend
	if frontendConfSrc == nil {
		frontendConfSrc = make(map[string]interface{})
	}

	frontendConf := make(map[string]interface{})
	utils.CopyStruct(&frontendConfSrc, &frontendConf)

	frontendConf["imgUpload"] = &imgUpload

	return Map{
		"img_upload":    imgUpload,
		"frontend_conf": frontendConf,
	}
}
