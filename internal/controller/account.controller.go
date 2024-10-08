package controller

import (
	"net/http"
	"simple_bank/internal/models"
	"simple_bank/internal/service"
	"simple_bank/package/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	CreateAccount(c *gin.Context)
	GetAccountByID(c *gin.Context)
	UpdateAccount(c *gin.Context)
	DeleteAccount(c *gin.Context)
	ListAccountsByUserID(c *gin.Context)
}

type accountController struct {
	accountService service.AccountService
}

func NewAccountController() AccountController {
	return &accountController{
		accountService: service.NewAccountService(),
	}
}

func (ac *accountController) CreateAccount(c *gin.Context) {
	var account models.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}

	err := ac.accountService.Create(c.Request.Context(), &account)
	if err != nil {

		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}

	response.SuccessResponse(c, 2000, account)
}

func (ac *accountController) GetAccountByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}

	account, err := ac.accountService.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorNotFoundResponse(c, 4001)
		return
	}

	response.SuccessResponse(c, 2000, account)
}

func (ac *accountController) UpdateAccount(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}

	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account.ID = uint(id)

	err = ac.accountService.Update(c.Request.Context(), &account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (ac *accountController) DeleteAccount(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}

	err = ac.accountService.Delete(c.Request.Context(), uint(id))
	if err != nil {
		response.ErrorBadGatewayResponse(c, 5000)
		return
	}

	response.NoContentResponse(c, 2040)
}

func (ac *accountController) ListAccountsByUserID(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}

	accounts, err := ac.accountService.ListByUserID(c.Request.Context(), uint(userID))
	if err != nil {
		response.ErrorBadGatewayResponse(c, 5000)
		return
	}

	response.SuccessResponse(c, 2000, accounts)
}
