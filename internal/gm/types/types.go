package types

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrNeedAuthorization       = errors.New("need authorization")
	ErrInvalidAuthorization    = errors.New("customer auth key is missing")
	ErrCustomerAlreadyExists   = errors.New("customer already registered")
	ErrRegisterParamsRequest   = errors.New("login and password required")
	ErrRegistration            = errors.New("internal error registration")
	ErrCustomerNotFound        = errors.New("customer not found")
	ErrCustomerLogin           = errors.New("customer invalid login")
	ErrInvalidOrderNumber      = errors.New("invalid order number")
	ErrOrderAlreadyExists      = errors.New("order number already exists")
	ErrOrderAnotherCustomer    = errors.New("order number used by another customer")
	ErrNotEnoughPoints         = errors.New("not enough points")
)

type Requester interface {
	GetCtx() *gin.Context
}

type ApiRequest struct {
	ctx *gin.Context `json:"-"`
}

type CustomerRegisterRequest struct {
	ApiRequest
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CustomerLoginRequest struct {
	ApiRequest
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AddOrderRequest struct {
	ApiRequest
	OrderNumber string
}

type CustomerWithdrawRequest struct {
	ApiRequest
	Order string  `json:"order"`
	Sum   float32 `json:"sum"`
}

type CustomerBalanceResponse struct {
	Balance  float32 `json:"balance"`
	Withdraw float32 `json:"withdraw"`
}

type ApiFunc func(Requester) (Response, error)

func NewRequest(ctx *gin.Context) ApiRequest {
	return ApiRequest{ctx: ctx}
}

func (r ApiRequest) GetCtx() *gin.Context {
	return r.ctx
}

type Order struct {
	Number     string  `db:"number" json:"number"`
	CustomerId int64   `db:"customer_id" json:"-"`
	Status     string  `db:"status" json:"status"`
	Accrual    float32 `db:"accrual" json:"accrual"`
	UploadedAt string  `db:"uploaded_at" json:"uploaded_at"`
}

type Withdraw struct {
	Number      string  `db:"number" json:"order"`
	CustomerId  int64   `db:"customer_id" json:"-"`
	Sum         float32 `db:"sum" json:"sum"`
	ProcessedAt string  `db:"processed_at" json:"processed_at"`
}

type Response any
