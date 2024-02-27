package models

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type DashBoardUser struct {
	TotalUsers  int `json:"Totaluser"`
	BlockedUser int `json:"Blockeduser"`
}

type DashBoardProduct struct {
	TotalProduct      int `json:"Totalproducts"`
	OutofStockProduct int `json:"Outofstock"`
}

type DashBoardOrder struct {
	CompletedOrder int
	PendingOrder   int
	CancelledOrder int
	TotalOrder     int
	TotalOrderItem int
}

type DashBoardRevenue struct {
	TodayRevenue float64
	MonthRevenue float64
	YearRevenue  float64
}

type DashBoardAmount struct {
	CreditedAmount float64
	PendingAmount  float64
}

type CompleteAdminDashboard struct {
	DashboardUser    DashBoardUser
	DashBoardProduct DashBoardProduct
	DashBoardOrder   DashBoardOrder
	DashBoardRevenue DashBoardRevenue
	DashBoardAmount  DashBoardAmount
}

type SalesReport struct {
	TotalSales      float64
	TotalOrders     int
	CompletedOrders int
	PendingOrders   int
	TrendingProduct string
}
