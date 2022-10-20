package handlers

import (
	"fmt"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/team"
	"github.com/ervera/tdlc-gin/pkg/web"
	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamService team.Service
}

func NewTeamHandler(t team.Service) *TeamHandler {
	return &TeamHandler{
		teamService: t,
	}
}

func (c *TeamHandler) CreateTeam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var team domain.Team
		err := ctx.ShouldBindJSON(&team)
		if err != nil {
			fmt.Println(err.Error())
			web.Error(ctx, 400, err.Error())
			return
		}
		resultTeam, err := c.teamService.Save(ctx, team)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Response(ctx, 200, resultTeam)
	}
}

func (c *TeamHandler) GetUserTeam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		team, err := c.teamService.GetAllByUserId(ctx)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Response(ctx, 200, team)
	}
}
