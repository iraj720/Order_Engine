package orderbook

const (
	Order_Status_UnSpecified = iota
	Order_Status_New
	Order_Status_Partially_Filled
	Order_Status_Filled
	Order_Status_Canceled
	Order_Status_Expired
	Order_Status_Rejected
	Order_Status_Partially_Cancelled
	Order_Status_Partially_Expired
	Order_Status_Semi_Expired
	Order_Status_Market_Cancelled

	Order_Action_Cancel

	Market_Action_Feed
	Market_Action_Close

	Market_Order
	Limit_Order
)
