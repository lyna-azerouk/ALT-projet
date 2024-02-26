package requests

const (
	SelectClientByEmailAndPasswordRequestTemplate  = "SELECT email, password, user_role FROM BL_USER WHERE email = $1 AND password = $2"
	SelectRestaurantByIdAndPasswordRequestTemplate = ""
	SelectMenusByRestaurantIdRequestTemplate       = "SELECT * FROM menus WHERE restaurant_id = $1"
	InsertNewClientRequestTemplate                 = ("INSERT into BL_USER (email, password, user_role) VALUES ($1, $2, $3)")
	InsertNewOrderRequestTemplate                  = ("INSERT into order_details (client_id, restaurant_id, price, order_status, order_date) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	InsertNewOrderItemRequestTemplate              = ("INSERT into order_items (order_id, menus_id, item_count) VALUES ($1, $2, $3)")
	SelectMenuByIdTemplate                         = "SELECT price FROM Menus where id= $1"
)
