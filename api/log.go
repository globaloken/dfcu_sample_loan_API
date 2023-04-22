package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	db "github.com/wizlif/dfcu_bank/db/sqlc"
	"github.com/wizlif/dfcu_bank/token"
)

func (server *Server) Log(ctx context.Context, t db.LogType, username *string) {
	u := ""
	if username != nil {
		u = *username
	}

	err := server.db.CreateLog(ctx, db.CreateLogParams{
		Type: t,
		Username: sql.NullString{
			String: u,
			Valid:  username != nil,
		},
	})

	if err != nil {
		log.Err(err).Msg("error creating log")
	}
}

func (server *Server) GetLogs(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	account, err := server.db.GetUser(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("no such user")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if account.Type != db.UserTypeADMIN {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("not authorized to access logs")))
		return
	}

	logs, err := server.db.GetLogStats(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, logs)
}
