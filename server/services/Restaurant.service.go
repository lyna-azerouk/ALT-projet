package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"serveur/server/const/requests"
	"serveur/server/database"
	"serveur/server/models"
)

const URL_TEMPLATE = "https://overpass-api.de/api/interpreter?data=%s"

func AllBouffluenceRestaurants() []models.RestaurantDetails {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	query := requests.SelectAllBouffluenceRestaurantsRequestTemplate
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var restaurants []models.RestaurantDetails
	for rows.Next() {
		var restaurant models.RestaurantDetails
		err := rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Affluence, &restaurant.OrderAverageDuration)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		restaurants = append(restaurants, restaurant)
	}

	return restaurants
}

/* RestaurantsAround return the list of restaurant around the given longitute and latitude
 * @param lon float64
 * @param lat float64
 * @param radius float64
 * @return List<OverPassRestaurant>
 */
func RestaurantsAround(lon float64, lat float64, radius float64) []models.OverPassRestaurant {
	query := fmt.Sprintf(`[out:json];
        node["amenity"="restaurant"](bbox:%f,%f,%f,%f);
        out;`, lat-radius, lon-radius, lat+radius, lon+radius)
	endpoint := fmt.Sprintf(URL_TEMPLATE, url.QueryEscape(query))
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	var overPassResponse models.OverPassResponse
	err = json.Unmarshal(body, &overPassResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return overPassResponse.Elements
}

func GetRestaurantDetails(restaurantID int) models.BouffluenceRestaurant {
	query := fmt.Sprintf("[out:json];node(%d);out;", restaurantID)
	endpoint := fmt.Sprintf(URL_TEMPLATE, url.QueryEscape(query))
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return models.BouffluenceRestaurant{}
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return models.BouffluenceRestaurant{}
	}
	var overPassResponse models.OverPassResponse
	err = json.Unmarshal(body, &overPassResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return models.BouffluenceRestaurant{}
	}
	// retrieve menu
	var menus = GetMenusByRestaurantId(restaurantID)
	restaurant := models.BouffluenceRestaurant{
		RestaurantDetails:    overPassResponse.Elements[0],
		Menu:                 menus,
		OrderAverageDuration: GetRestaurantOrderAverageDuration(restaurantID),
	}
	return restaurant

}

func GetRestaurantOrderAverageDuration(restaurantID int) int {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	query := requests.SelectRestaurantOrderAverageDurationRequestTemplate
	row := db.QueryRow(query, restaurantID)
	var orderAverageDuration int
	err = row.Scan(&orderAverageDuration)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return orderAverageDuration
}

func GetMenusByRestaurantId(restaurantId int) []models.Menu {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	query := requests.SelectMenusByRestaurantIdRequestTemplate
	rows, err := db.Query(query, restaurantId)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	var menus []models.Menu
	for rows.Next() {
		var menu models.Menu
		var price []uint8
		err := rows.Scan(&menu.Id, &menu.Name, &menu.RestaurantID, &price, &menu.Description, &menu.Image)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		menu.Price = float64(price[0])
		menus = append(menus, menu)
	}
	return menus
}

func AddMenuItem(restaurantId int, menuItem models.MenuItem) error {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	query := requests.InsertMenuItemRequestTemplate
	_, err = db.Exec(query, restaurantId, menuItem.Name, menuItem.Price, menuItem.Description, menuItem.Image)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
