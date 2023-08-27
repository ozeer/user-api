package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"user-api/forms"
	"user-api/global"
	"user-api/global/response"
	"user-api/proto"
	"user-api/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func GetUserList(ctx *gin.Context) {
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.Conf.Server.Host, global.Conf.Server.Port), grpc.WithInsecure())

	if err != nil {
		global.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	userClient := proto.NewUserClient(userConn)
	page := ctx.DefaultQuery("page", utils.Uint32ToString(global.Page))
	pageInt := utils.StringToUint32(page)
	pageSize := ctx.DefaultQuery("size", utils.Uint32ToString(global.PageSize))
	pageSizeInt := utils.StringToUint32(pageSize)

	resp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Page: pageInt,
		Size: pageSizeInt,
	})

	if err != nil {
		global.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)

	for _, v := range resp.Data {
		// data := make(map[string]interface{}, 0)

		user := response.UserResponse{
			Id:       v.Id,
			NickName: v.Nickname,
			// Birthday: time.Time(time.Unix(int64(v.Birthday), 0)),
			// Birthday: time.Time(time.Unix(int64(v.Birthday), 0)).Format("2006-01-02"),
			Birthday: response.JsonTime(time.Unix(int64(v.Birthday), 0)),

			Gender: v.Gender,
			Mobile: v.Mobile,
		}

		result = append(result, user)
	}

	ctx.JSON(http.StatusOK, result)
}

func PasswordLogin(c *gin.Context) {
	// 表单验证
	PasswordLoginForm := forms.PasswordLoginForm{}

	if err := c.ShouldBindJSON(&PasswordLoginForm); err != nil {
		global.HandleValidatorError(err, c)
		return
	}
}
