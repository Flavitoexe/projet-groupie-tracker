package controllers

import (
	"fmt"
	"guide/helper"
	"guide/services"
	"log"
	"net/http"
	"slices"
	"strings"
)

func DisplayListCountries(w http.ResponseWriter, r *http.Request) {
	data, statusCode, err := services.GetAllCountries()
	if statusCode != http.StatusOK || err != nil {
		http.Redirect(w, r, "/error", statusCode)
		// http.Error(w, fmt.Sprintf("Erreur service - code : %d \n message: %s", statusCode, err.Error()), statusCode)
		return
	}

	helper.RenderTemplate(w, r, "menu", data)
}

func DisplayCountriesByRegion(w http.ResponseWriter, r *http.Request) {

	region := r.FormValue("region")
	fmt.Println(region)

	data, dataStatusCode, dataError := services.GetAllCountries()
	if dataStatusCode != http.StatusOK || dataError != nil {
		log.Printf("Erreur DisplayCountriesByRegion - %s", dataError.Error())
		http.Error(w, dataError.Error(), http.StatusMovedPermanently)
		return
	}

	validCountries := []services.Country{}

	for _, item := range *data {
		if strings.ToLower(item.Region) == region {
			validCountries = append(validCountries, item)
		}
	}

	helper.RenderTemplate(w, r, "region_choice", validCountries)
}

func DisplayCountryDetails(w http.ResponseWriter, r *http.Request) {

	alpha3Code := r.FormValue("alpha3code")
	fmt.Println(alpha3Code)

	data, dataStatusCode, dataError := services.GetCountryByAlpha3(alpha3Code)
	if dataStatusCode != http.StatusOK || dataError != nil {
		log.Printf("Erreur DisplayCountriesDeatils - %s", dataError.Error())
		http.Error(w, dataError.Error(), http.StatusMovedPermanently)
		return
	}

	helper.RenderTemplate(w, r, "country_details", data)
}

func DisplayError(w http.ResponseWriter, r *http.Request) {

	helper.RenderTemplate(w, r, "error", nil)
}

type TempCountryData struct {
	Data           *[]services.Country
	DataStatusCode int
	DataErr        error
}

func DisplayCountriesByName(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	fmt.Println(name)

	var dataTemp TempCountryData

	dataTemp.Data, dataTemp.DataStatusCode, dataTemp.DataErr = services.GetCountryByName(name)
	if dataTemp.DataStatusCode != http.StatusOK && dataTemp.DataStatusCode != http.StatusNotFound {
		log.Printf("Erreur DisplayCountriesByName - %s", dataTemp.DataErr.Error())
		http.Redirect(w, r, "/error", dataTemp.DataStatusCode)
		// http.Error(w, dataTemp.DataErr.Error(), http.StatusMovedPermanently)
		return
	}

	helper.RenderTemplate(w, r, "country_choice", dataTemp)
}

func DisplayCountriesByLang(w http.ResponseWriter, r *http.Request) {

	lang := r.URL.Query()["lang"]
	fmt.Println(lang)

	var dataTemp TempCountryData

	dataTemp.Data, dataTemp.DataStatusCode, dataTemp.DataErr = services.GetAllCountries()
	if dataTemp.DataStatusCode != http.StatusOK && dataTemp.DataStatusCode != http.StatusNotFound {
		log.Printf("Erreur DisplayCountriesByLang - %s", dataTemp.DataErr.Error())
		http.Redirect(w, r, "/error", dataTemp.DataStatusCode)
		// http.Error(w, dataTemp.DataErr.Error(), http.StatusMovedPermanently)
		return
	}

	validCountries := []services.Country{}

	for index := 0; index < len(*dataTemp.Data); index++ {

		languages := []string{}
		for _, lang := range (*dataTemp.Data)[index].Languages {
			languages = append(languages, strings.ToLower(lang.Name))
		}

		checkLang := slices.ContainsFunc(languages, func(item string) bool {
			return slices.Contains(lang, item)
		})

		if checkLang || len(lang) == 0 {
			validCountries = append(validCountries, (*dataTemp.Data)[index])
		}
	}

	*dataTemp.Data = validCountries
	println(*dataTemp.Data)

	helper.RenderTemplate(w, r, "country_lang", dataTemp)
}
