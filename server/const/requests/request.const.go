package requests

const (
	SelectClientByEmailAndPasswordRequestTemplate  = "SELECT * FROM BL_USER WHERE email = $1 AND password = $2"
	SelectRestaurantByIdAndPasswordRequestTemplate = ""
	SelectMenusByRestaurantIdRequestTemplate       = "SELECT * FROM menus WHERE restaurant_id = $1"
	InsertNewClientRequestTemplate                 = ("INSERT into BL_USER (email, password, user_role) VALUES ($1, $2, $3)")
)
