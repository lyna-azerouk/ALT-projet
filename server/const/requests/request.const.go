package requests

const (
	SelectClientByEmailAndPasswordRequestTemplate       = "SELECT id, email, user_role FROM BL_USER WHERE email = $1 AND password = $2"
	SelectRestaurantByIdAndPasswordRequestTemplate      = "SELECT * FROM restaurant WHERE id = $1 AND password = $2"
	SelectMenusByRestaurantIdRequestTemplate            = "SELECT * FROM menus WHERE restaurant_id = $1"
	InsertNewClientRequestTemplate                      = ("INSERT into BL_USER (email, password, user_role, first_name, last_name) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	InsertNewOrderRequestTemplate                       = ("INSERT into order_details (client_id, restaurant_id, price, order_status, order_date) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	InsertNewOrderItemRequestTemplate                   = ("INSERT into order_items (order_id, menus_id, item_count) VALUES ($1, $2, $3)")
	SelectMenuByIdTemplate                              = "SELECT price FROM Menus where id= $1"
	UpdateStatusOrderRequestTemplate                    = "UPDATE order_details SET order_status = $1 WHERE id = $2;"
	GetOrderRequestTemplate                             = "Select * from order_details where id = $1"
	GetOrderItemsRequestTemplate                        = "Select menus_id, item_count FROM order_items where order_id =$1"
	DeleteOrderItemsRequestTemplate                     = "DELETE FROM order_items WHERE order_id = $1;"
	DeleteOrderDetailsRequestTemplate                   = "DELETE FROM order_details WHERE id = $1"
	SelectRestaurantOrderAverageDurationRequestTemplate = "SELECT order_average_duration FROM restaurant WHERE id = $1"
	InsertCodeRequestTemplate                           = "INSERT into Orders_Code (order_id, code_confirmation) VALUES ($1, $2)"
	GetOrderCodeTemplate                                = "Select code_confirmation FROM Orders_Code WHERE order_id= $1 AND code_confirmation=$2 "
	GetUserDetailsRequestTemplate                       = "SELECT id, first_name, last_name, email FROM BL_USER WHERE id = $1"
	GetUserOrdersTemplate                               = "Select id From order_details where client_id= $1"
	GetRestaurantOrdersTemplate						 = "Select id From order_details where restaurant_id= $1"
)
