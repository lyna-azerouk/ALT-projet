package requests

const (
	SelectClientByEmailAndPasswordRequestTemplate  = "SELECT * FROM BL_USER WHERE email = $1 AND password = $2"
	SelectRestaurantByIdAndPasswordRequestTemplate = ""
)
