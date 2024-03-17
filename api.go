package merchantapiinterfaces

import "github.com/gin-gonic/gin"

func MerchantApiBinding(router *gin.Engine, api MerchantApi) {
	router.GET("/merchants/self", api.GetSelfMerchant)
	router.GET("/merchants", api.ListMerchants)
	router.POST("/merchants", api.CreateMerchant)
	router.GET("/merchants/:merchantId", api.GetMerchant)
	router.PUT("/merchants/:merchantId", api.UpdateMerchant)
	router.DELETE("/merchants/:merchantId", api.DeleteMerchant)
	router.POST("/merchants/:merchantId/contact-form", api.SendContactForm)
	router.GET("/merchants/self/config", api.GetSelfMerchantConfig)
	router.PUT("/merchants/self/config", api.UpdateSelfMerchantConfig)
	router.GET("/merchants/self/contact-form", api.ListSelfContactForm)
	router.GET("/merchants/self/members", api.ListSelfMembers)
	router.GET("/merchants/self/contact-form/:id", api.GetSelfContactForm)
	router.DELETE("/merchants/self/contact-form/:id", api.DeleteSelfContactForm)
	router.POST("/merchants/:merchantId/verification-code", api.SendMerchantVerificationCode)
	router.PUT("/merchants/:merchantId/password", api.UpdateMerchantPassword)
}
func SpaceApiBinding(router *gin.Engine, api SpaceApi) {
	router.GET("/spaces/:spaceId", api.GetSpace)
	router.PUT("/spaces/:spaceId", api.UpdateSpace)
	router.DELETE("/spaces/:spaceId", api.DeleteSpace)
	router.POST("/spaces", api.CreateSpace)
	router.GET("/spaces", api.ListSpaces)
	router.GET("/spaces/:spaceId/parent", api.GetSpaceParent)
	router.GET("/spaces/:spaceId/children", api.ListSpaceChildren)
	router.GET("/spaces/printers/:printerId", api.GetPrinter)
	router.PUT("/spaces/printers/:printerId", api.UpdatePrinter)
	router.DELETE("/spaces/printers/:printerId", api.DeletePrinter)
	router.GET("/spaces/:spaceId/printers", api.ListPrinters)
	router.POST("/spaces/:spaceId/printers", api.CreatePrinter)
}
func OrderApiBinding(router *gin.Engine, api OrderApi) {
	router.GET("/orders/bills/:billId", api.GetBill)
	router.PUT("/spaces/tables/:id", api.UpdateTable)
	router.DELETE("/spaces/tables/:id", api.DeleteTable)
	router.PUT("/orders/:orderId/tableLabel", api.UpdateOrderTableLabel)
	router.PATCH("/orders/:orderId/addition", api.AddOrder)
	router.POST("/orders/bills/print", api.PrintBill)
	router.PUT("/spaces/:spaceId/foods/categories", api.UpdateFoodCategories)
	router.GET("/spaces/:spaceId/foods/categories", api.ListFoodCategories)
	router.GET("/spaces/foods/:id/printers", api.ListFoodPrinters)
	router.PUT("/spaces/foods/:id/printers", api.UpdateFoodPrinters)
	router.GET("/spaces/:spaceId/tables", api.ListTables)
	router.POST("/spaces/:spaceId/tables", api.CreateTable)
	router.GET("/spaces/:spaceId/foods", api.ListFoods)
	router.POST("/spaces/:spaceId/foods", api.CreateFood)
	router.PATCH("/orders/:orderId/cancel", api.CancelOrder)
	router.GET("/orders/bills/statistics", api.GetBillStatistics)
	router.POST("/orders/bills", api.CreateBill)
	router.GET("/orders/bills", api.ListBills)
	router.POST("/spaces/:spaceId/orders", api.CreateOrder)
	router.GET("/spaces/:spaceId/orders", api.ListOrders)
	router.DELETE("/spaces/foods/:id", api.DeleteFood)
	router.GET("/spaces/foods/:id", api.GetFoodById)
	router.PUT("/spaces/foods/:id", api.UpdateFood)
	router.PUT("/spaces/foods/:id/image", api.UpdateFoodImage)
}
func MemberApiBinding(router *gin.Engine, api MemberApi) {
	router.POST("/merchants/:merchantId/members/signup", api.SignupMember)
	router.GET("/merchants/:merchantId/members", api.ListMembers)
	router.POST("/merchants/:merchantId/members", api.CreateMember)
}
func SessionApiBinding(router *gin.Engine, api SessionApi) {
	router.POST("/sessions", api.CreateSession)
}
func VerificationApiBinding(router *gin.Engine, api VerificationApi) {
	router.POST("/verification", api.SendVerificationCode)
}

type MerchantApi interface {
	GetSelfMerchant(ctx *gin.Context)
	ListMerchants(ctx *gin.Context)
	CreateMerchant(ctx *gin.Context)
	GetMerchant(ctx *gin.Context)
	UpdateMerchant(ctx *gin.Context)
	DeleteMerchant(ctx *gin.Context)
	SendContactForm(ctx *gin.Context)
	GetSelfMerchantConfig(ctx *gin.Context)
	UpdateSelfMerchantConfig(ctx *gin.Context)
	ListSelfContactForm(ctx *gin.Context)
	ListSelfMembers(ctx *gin.Context)
	GetSelfContactForm(ctx *gin.Context)
	DeleteSelfContactForm(ctx *gin.Context)
	SendMerchantVerificationCode(ctx *gin.Context)
	UpdateMerchantPassword(ctx *gin.Context)
}
type SpaceApi interface {
	GetSpace(ctx *gin.Context)
	UpdateSpace(ctx *gin.Context)
	DeleteSpace(ctx *gin.Context)
	CreateSpace(ctx *gin.Context)
	ListSpaces(ctx *gin.Context)
	GetSpaceParent(ctx *gin.Context)
	ListSpaceChildren(ctx *gin.Context)
	GetPrinter(ctx *gin.Context)
	UpdatePrinter(ctx *gin.Context)
	DeletePrinter(ctx *gin.Context)
	ListPrinters(ctx *gin.Context)
	CreatePrinter(ctx *gin.Context)
}
type OrderApi interface {
	GetBill(ctx *gin.Context)
	UpdateTable(ctx *gin.Context)
	DeleteTable(ctx *gin.Context)
	UpdateOrderTableLabel(ctx *gin.Context)
	AddOrder(ctx *gin.Context)
	PrintBill(ctx *gin.Context)
	UpdateFoodCategories(ctx *gin.Context)
	ListFoodCategories(ctx *gin.Context)
	ListFoodPrinters(ctx *gin.Context)
	UpdateFoodPrinters(ctx *gin.Context)
	ListTables(ctx *gin.Context)
	CreateTable(ctx *gin.Context)
	ListFoods(ctx *gin.Context)
	CreateFood(ctx *gin.Context)
	CancelOrder(ctx *gin.Context)
	GetBillStatistics(ctx *gin.Context)
	CreateBill(ctx *gin.Context)
	ListBills(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	ListOrders(ctx *gin.Context)
	DeleteFood(ctx *gin.Context)
	GetFoodById(ctx *gin.Context)
	UpdateFood(ctx *gin.Context)
	UpdateFoodImage(ctx *gin.Context)
}
type MemberApi interface {
	SignupMember(ctx *gin.Context)
	ListMembers(ctx *gin.Context)
	CreateMember(ctx *gin.Context)
}
type SessionApi interface {
	CreateSession(ctx *gin.Context)
}
type VerificationApi interface {
	SendVerificationCode(ctx *gin.Context)
}
type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"

type Spec struct {
	Attribute string `json:"attribute" xml:"attribute"`
	Optioned  string `json:"optioned" xml:"optioned"`
}
type PrinterModel string

const F80MM PrinterModel = "F80MM"
const F58MM PrinterModel = "F58MM"

type Table struct {
	Id    string `json:"id" xml:"id"`
	Label string `json:"label" xml:"label"`
}
type MerchantList struct {
	Items      []Merchant `json:"items" xml:"items"`
	Pagination Pagination `json:"pagination" xml:"pagination"`
}
type CreateSessionRequest struct {
	ShortMerchantId *string `json:"shortMerchantId" xml:"shortMerchantId"`
	Account         *string `json:"account" xml:"account"`
	Password        *string `json:"password" xml:"password"`
}
type Role string

const ROOT Role = "ROOT"
const ADMIN Role = "ADMIN"

type CreateOrderRequest struct {
	Remark     *string    `json:"remark" xml:"remark"`
	Foods      []FoodSpec `json:"foods" xml:"foods"`
	TableLabel *string    `json:"tableLabel" xml:"tableLabel"`
}
type UpdateFoodPrintersRequest struct {
	PrinterIds []string `json:"printerIds" xml:"printerIds"`
}
type SaveFoodRequest struct {
	FixedOffset *int64          `json:"fixedOffset" xml:"fixedOffset"`
	Image       string          `json:"image" xml:"image"`
	Categories  []string        `json:"categories" xml:"categories"`
	Status      *FoodStatus     `json:"status" xml:"status"`
	Attributes  []FoodAttribute `json:"attributes" xml:"attributes"`
	Name        string          `json:"name" xml:"name"`
	Description string          `json:"description" xml:"description"`
	Price       int64           `json:"price" xml:"price"`
}
type FoodAttributesOption struct {
	Label string `json:"label" xml:"label"`
	Extra *int64 `json:"extra" xml:"extra"`
}
type AddOrderRequest struct {
	Foods []FoodSpec `json:"foods" xml:"foods"`
}
type CancelOrderRequest struct {
	Foods []FoodSpec `json:"foods" xml:"foods"`
}
type Order struct {
	Foods      []FoodSpec  `json:"foods" xml:"foods"`
	Status     OrderStatus `json:"status" xml:"status"`
	CreatedAt  int64       `json:"createdAt" xml:"createdAt"`
	UpdatedAt  int64       `json:"updatedAt" xml:"updatedAt"`
	Id         string      `json:"id" xml:"id"`
	Code       string      `json:"code" xml:"code"`
	TableLabel *string     `json:"tableLabel" xml:"tableLabel"`
}
type Pagination struct {
	Index int64 `json:"index" xml:"index"`
	Limit int64 `json:"limit" xml:"limit"`
	Total int64 `json:"total" xml:"total"`
}
type Bill struct {
	Id        string  `json:"id" xml:"id"`
	Amount    int64   `json:"amount" xml:"amount"`
	Orders    []Order `json:"orders" xml:"orders"`
	CreatedAt int64   `json:"createdAt" xml:"createdAt"`
}
type OrderStatus string

const SUBMITTED OrderStatus = "SUBMITTED"
const CANCELLED OrderStatus = "CANCELLED"
const COMPLETED OrderStatus = "COMPLETED"

type Space struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
	Description string `json:"description" xml:"description"`
}
type CreateVerificationCodeRequest struct {
	MerchantId  *string      `json:"merchantId" xml:"merchantId"`
	PhoneNumber *PhoneNumber `json:"phoneNumber" xml:"phoneNumber"`
}
type CreateMerchantRequest struct {
	ShortMerchantId string `json:"shortMerchantId" xml:"shortMerchantId"`
	Account         string `json:"account" xml:"account"`
	Password        string `json:"password" xml:"password"`
}
type ContactForm struct {
	Phone   *PhoneNumber `json:"phone" xml:"phone"`
	Email   *string      `json:"email" xml:"email"`
	Subject *string      `json:"subject" xml:"subject"`
	Message *string      `json:"message" xml:"message"`
	Name    *string      `json:"name" xml:"name"`
}
type Printer struct {
	Sn    string       `json:"sn" xml:"sn"`
	Type  PrinterType  `json:"type" xml:"type"`
	Model PrinterModel `json:"model" xml:"model"`
	Id    string       `json:"id" xml:"id"`
	Name  string       `json:"name" xml:"name"`
}
type FoodAttribute struct {
	Label   string                 `json:"label" xml:"label"`
	Options []FoodAttributesOption `json:"options" xml:"options"`
}
type UpdateOrderTableLabelRequest struct {
	TableLabel *string `json:"tableLabel" xml:"tableLabel"`
}
type SavePrinter struct {
	Name  string       `json:"name" xml:"name"`
	Sn    string       `json:"sn" xml:"sn"`
	Type  PrinterType  `json:"type" xml:"type"`
	Model PrinterModel `json:"model" xml:"model"`
}
type MerchantConfig struct {
	Currency *Currency `json:"currency" xml:"currency"`
}
type SaveSpaceRequest struct {
	Name        string  `json:"name" xml:"name"`
	Description string  `json:"description" xml:"description"`
	ParentId    *string `json:"parentId" xml:"parentId"`
}
type PhoneNumber struct {
	CountryCode string `json:"countryCode" xml:"countryCode"`
	Number      string `json:"number" xml:"number"`
}
type BillStatistics struct {
	Count  int64 `json:"count" xml:"count"`
	Amount int64 `json:"amount" xml:"amount"`
}
type FoodStatus string

const AVAILABLE FoodStatus = "AVAILABLE"
const UNAVAILABLE FoodStatus = "UNAVAILABLE"

type SaveTableRequest struct {
	Label string `json:"label" xml:"label"`
}
type ContactFormList struct {
	Items      []ContactForm `json:"items" xml:"items"`
	Pagination Pagination    `json:"pagination" xml:"pagination"`
}
type CreateMemberRequest struct {
	MerchantId       *string      `json:"merchantId" xml:"merchantId"`
	VerificationCode *string      `json:"verificationCode" xml:"verificationCode"`
	PhoneNumber      *PhoneNumber `json:"phoneNumber" xml:"phoneNumber"`
}
type MemberList struct {
	Pagination Pagination `json:"pagination" xml:"pagination"`
	Items      []Member   `json:"items" xml:"items"`
}
type CreateAdminRequest struct {
	Account  string `json:"account" xml:"account"`
	Password string `json:"password" xml:"password"`
	Role     *Role  `json:"role" xml:"role"`
}
type Session struct {
	Token string `json:"token" xml:"token"`
}
type UpdatePasswordRequest struct {
	Password    string  `json:"password" xml:"password"`
	OldPassword *string `json:"oldPassword" xml:"oldPassword"`
}
type Merchant struct {
	UpdatedAt       int64  `json:"updatedAt" xml:"updatedAt"`
	Id              string `json:"id" xml:"id"`
	ShortMerchantId string `json:"shortMerchantId" xml:"shortMerchantId"`
	Account         string `json:"account" xml:"account"`
	Name            string `json:"name" xml:"name"`
	Description     string `json:"description" xml:"description"`
	CreatedAt       int64  `json:"createdAt" xml:"createdAt"`
}
type BillList struct {
	Items      []Bill     `json:"items" xml:"items"`
	Pagination Pagination `json:"pagination" xml:"pagination"`
}
type CreateBillRequest struct {
	Amount   int64    `json:"amount" xml:"amount"`
	OrderIds []string `json:"orderIds" xml:"orderIds"`
}
type Currency string

const HKD Currency = "HKD"
const USD Currency = "USD"
const CNY Currency = "CNY"
const MOP Currency = "MOP"

type Menu struct {
	Foods []Food `json:"foods" xml:"foods"`
}
type SpaceList struct {
	Items      []Space    `json:"items" xml:"items"`
	Pagination Pagination `json:"pagination" xml:"pagination"`
}
type SignupMemberRequest struct {
	MerchantId       string      `json:"merchantId" xml:"merchantId"`
	PhoneNumber      PhoneNumber `json:"phoneNumber" xml:"phoneNumber"`
	VerificationCode string      `json:"verificationCode" xml:"verificationCode"`
}
type Member struct {
	Id          string      `json:"id" xml:"id"`
	MerchantId  string      `json:"merchantId" xml:"merchantId"`
	Name        string      `json:"name" xml:"name"`
	PhoneNumber PhoneNumber `json:"phoneNumber" xml:"phoneNumber"`
	CreatedAt   int64       `json:"createdAt" xml:"createdAt"`
	UpdatedAt   int64       `json:"updatedAt" xml:"updatedAt"`
}
type Food struct {
	Image       string          `json:"image" xml:"image"`
	Categories  []string        `json:"categories" xml:"categories"`
	UpdatedAt   int64           `json:"updatedAt" xml:"updatedAt"`
	Id          string          `json:"id" xml:"id"`
	Name        string          `json:"name" xml:"name"`
	Description string          `json:"description" xml:"description"`
	Price       int64           `json:"price" xml:"price"`
	FixedOffset *int64          `json:"fixedOffset" xml:"fixedOffset"`
	Status      FoodStatus      `json:"status" xml:"status"`
	Attributes  []FoodAttribute `json:"attributes" xml:"attributes"`
	CreatedAt   int64           `json:"createdAt" xml:"createdAt"`
}
type FoodSpec struct {
	Food Food    `json:"food" xml:"food"`
	Spec *[]Spec `json:"spec" xml:"spec"`
}
type PrinterType string

const KITCHEN PrinterType = "KITCHEN"
const CASHIER PrinterType = "CASHIER"
