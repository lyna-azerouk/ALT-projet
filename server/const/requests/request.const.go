package requests

const USER_TABLE = "bl_users"
const RESTAURANT_TABLE = "restaurant"
const MENU_TABLE = "menus"
const ORDER_TABLE = "order_details"
const ORDER_ITEMS_TABLE = "order_items"
const ORDER_CODE_TABLE = "order_code"

const (
	SelectClientByEmailAndPasswordRequestTemplate       = "SELECT id, email, user_role FROM " + USER_TABLE + " WHERE email = $1 AND password = $2"
	SelectRestaurantByIdAndPasswordRequestTemplate      = "SELECT * FROM " + RESTAURANT_TABLE + " WHERE id = $1 AND password = $2"
	SelectMenusByRestaurantIdRequestTemplate            = "SELECT * FROM " + MENU_TABLE + " WHERE restaurant_id = $1"
	InsertNewClientRequestTemplate                      = ("INSERT into " + USER_TABLE + " (email, password, user_role, last_name, first_name) VALUES ($1, $2, $3, $4, $5)")
	InsertNewOrderRequestTemplate                       = ("INSERT into " + ORDER_TABLE + " (client_id, restaurant_id, price, order_status, order_date) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	InsertNewOrderItemRequestTemplate                   = ("INSERT into " + ORDER_ITEMS_TABLE + " (order_id, menus_id, item_count) VALUES ($1, $2, $3)")
	SelectMenuByIdTemplate                              = "SELECT price FROM " + MENU_TABLE + " where id= $1"
	UpdateStatusOrderRequestTemplate                    = "UPDATE " + ORDER_TABLE + " SET order_status = $1 WHERE id = $2;"
	GetOrderRequestTemplate                             = "Select * from " + ORDER_TABLE + " where id = $1"
	GetOrderItemsRequestTemplate                        = "Select menus_id, item_count FROM " + ORDER_ITEMS_TABLE + " where order_id =$1"
	DeleteOrderItemsRequestTemplate                     = "DELETE FROM " + ORDER_ITEMS_TABLE + " WHERE order_id = $1;"
	DeleteOrderDetailsRequestTemplate                   = "DELETE FROM " + ORDER_TABLE + " WHERE id = $1"
	SelectRestaurantOrderAverageDurationRequestTemplate = "SELECT order_average_duration FROM " + RESTAURANT_TABLE + " WHERE id = $1"
	InsertCodeRequestTemplate                           = "INSERT into " + ORDER_CODE_TABLE + " (order_id, code_confirmation) VALUES ($1, $2)"
	GetOrderCodeTemplate                                = "Select code_confirmation FROM " + ORDER_CODE_TABLE + " WHERE order_id= $1 AND code_confirmation=$2 "
	GetUserDetailsRequestTemplate                       = "SELECT id, first_name, last_name, email FROM " + USER_TABLE + " WHERE id = $1"
	GetUserOrdersTemplate                               = "Select id From " + ORDER_TABLE + " where client_id= $1"
	GetRestaurantOrdersTemplate                         = "Select id From " + ORDER_TABLE + " where restaurant_id= $1"
	GetRestaurantOrdersDetailsTemplate                  = "SELECT OD.id, OD.client_id, OD.restaurant_id, OD.order_status, OD.price, OD.order_date, OI.menus_id, OI.item_count FROM " + ORDER_TABLE + " OD LEFT JOIN " + ORDER_ITEMS_TABLE + " OI ON OI.order_id = OD.id WHERE OD.restaurant_id=$1"
	UpdateAffluenceForRestaurantVoteRequestTemplate     = "UPDATE " + RESTAURANT_TABLE + " SET affluence = $1 WHERE id = $2 RETURNING affluence"
	SelectAffluenceRequestTemplate                      = "SELECT affluence FROM " + RESTAURANT_TABLE + " WHERE id = $1"
)
