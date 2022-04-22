package http

import (
	"github.com/ArtalkJS/ArtalkGo/model"
	"github.com/labstack/echo/v4"
)

type ParamsAdminSiteDel struct {
	ID         uint `mapstructure:"id" param:"required"`
	DelContent bool `mapstructure:"del_content"`
}

func (a *action) AdminSiteDel(c echo.Context) error {
	if isOK, resp := AdminOnly(c); !isOK {
		return resp
	}

	var p ParamsAdminSiteDel
	if isOK, resp := ParamsDecode(c, ParamsAdminSiteDel{}, &p); !isOK {
		return resp
	}

	site := model.FindSiteByID(p.ID)
	if site.IsEmpty() {
		return RespError(c, "site 不存在")
	}

	err := model.DelSite(&site, !p.DelContent)
	if err != nil {
		return RespError(c, "site 删除失败")
	}

	return RespSuccess(c)
}
