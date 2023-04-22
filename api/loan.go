package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/wizlif/dfcu_bank/db/sqlc"
	"github.com/wizlif/dfcu_bank/token"
)

type CreateLoanRequest struct {
	Amount int64 `json:"amount" binding:"required,gt=0"`
}

func (server *Server) CreateLoan(ctx *gin.Context) {
	var req CreateLoanRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	loan, err := server.db.CreateLoan(ctx, db.CreateLoanParams{
		Username: authPayload.Username,
		Amount:   req.Amount,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, loan)
}

type getAccountLoanParams struct {
	AccountNo string `uri:"acc_no" binding:"required,min=10,max=10"`
}

func (server *Server) GetLoanAccount(ctx *gin.Context) {
	var req getAccountLoanParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		go server.Log(ctx, db.LogTypeFAILEDVALIDATION, nil)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	account, err := server.db.GetUserByAccNo(ctx, req.AccountNo)
	if err != nil {
		if err == sql.ErrNoRows {
			go server.Log(ctx, db.LogTypeFAILEDVALIDATION, &authPayload.Username)
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New("no such account")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if account.Username != authPayload.Username && account.Type != db.UserTypeADMIN {
		go server.Log(ctx, db.LogTypeFAILEDVALIDATION, &authPayload.Username)
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("not authorized to acces this account")))
		return
	}

	loans, err := server.db.GetLoanByAccUsername(ctx, authPayload.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if len(loans) == 0 {
		go server.Log(ctx, db.LogTypeNEGATIVEREQUEST, &authPayload.Username)
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("no loans")))
		return
	}

	go server.Log(ctx, db.LogTypePOSITIVEREQUEST, &authPayload.Username)
	ctx.JSON(http.StatusOK, loans)
}
