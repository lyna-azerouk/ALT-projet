package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"serveur/server/models"
)

const URL_TEMPLATE = "https://overpass-api.de/api/interpreter?data=%s"

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
	println(string(body))
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
