package routes

import (
	"guide/controllers"
	"net/http"
)

func geoFactsRoutes(router *http.ServeMux) {

	router.HandleFunc("/", controllers.DisplayListCountries)

	router.HandleFunc("/countries", controllers.DisplayCountriesByRegion)

	router.HandleFunc("/country/details", controllers.DisplayCountryDetails)

	router.HandleFunc("/country", controllers.DisplayCountriesByName)

	router.HandleFunc("/countries-by-lang", controllers.DisplayCountriesByLang)

	router.HandleFunc("/error", controllers.DisplayError)

}
