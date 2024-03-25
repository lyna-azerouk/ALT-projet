package database

const PROFILE = "dev" // dev or prod

const (
	HOST_DEV     = "sour-yak-9161.7tc.aws-eu-central-1.cockroachlabs.cloud"
	PORT_DEV     = 26257
	USER_DEV     = "bouffluence"
	PASSWORD_DEV = "b7kYcKKNkWPqwqYTAj6NKg"
	DBNAME_DEV   = "bouffluence_db"
)

const (
	HOST_PROD     = "sour-yak-9161.7tc.aws-eu-central-1.cockroachlabs.cloud"
	PORT_PROD     = 26257
	USER_PROD     = "bouffluence"
	PASSWORD_PROD = "b7kYcKKNkWPqwqYTAj6NKg"
	DB_NAME_PROD  = "bouffluence_db"
)

func GetHost() string {
	if PROFILE == "dev" {
		return HOST_DEV
	}
	return HOST_PROD
}

func GetPort() int {
	if PROFILE == "dev" {
		return PORT_DEV
	}
	return PORT_PROD
}

/** Return the user name for the database connection *
 * according to the profile
 * @return string
 */
func GetUser() string {
	if PROFILE == "dev" {
		return USER_DEV
	}
	return USER_PROD
}

/** Return the password for the database connection *
 * according to the profile
 * @return string
 */
func GetPassword() string {
	if PROFILE == "dev" {
		return PASSWORD_DEV
	}
	return PASSWORD_PROD
}

/** Return the database name for the database connection *
 * according to the profile
 * @return string
 */
func GetDBName() string {
	if PROFILE == "dev" {
		return DBNAME_DEV
	}
	return DB_NAME_PROD
}