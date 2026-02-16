package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Country API Countries
type Country struct {
	Name           string   `json:"name"`
	Alpha2Code     string   `json:"alpha2Code"`
	Alpha3Code     string   `json:"alpha3Code"`
	Capital        string   `json:"capital"`
	Region         string   `json:"region"`
	Subregion      string   `json:"subregion"`
	Population     int      `json:"population"`
	Area           float64  `json:"area"`
	TopLevelDomain []string `json:"topLevelDomain"`
	Timezones      []string `json:"timezones"`
	Borders        []string `json:"borders"`
	Flags          struct {
		Png string `json:"png"`
	} `json:"flags"`
	Currencies []struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Languages []Languages `json:"languages"`
}

type Languages struct {
	Name string `json:"name"`
}

func GetAllCountries() (*[]Country, int, error) {

	_client := http.Client{Timeout: time.Second * 5}

	request, requestErr := http.NewRequest(http.MethodGet, "https://www.apicountries.com/countries", nil)
	if requestErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur préparation requête - %s", requestErr.Error())
	}

	response, responseErr := _client.Do(request)
	if responseErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur envoi requête - %s", responseErr.Error())
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil,
			response.StatusCode,
			fmt.Errorf("Erreur réponse - %s", response.Status)
	}

	var countriesList []Country

	decodeErr := json.NewDecoder(response.Body).Decode(&countriesList)
	if decodeErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur decode données - %s", decodeErr.Error())
	}

	return &countriesList, response.StatusCode, nil
}

func GetCountriesByRegion(region string) (*[]Country, int, error) {

	_client := http.Client{Timeout: time.Second * 5}

	request, requestError := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.apicountries.com/region/%s", region), nil)
	if requestError != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur préparation requête - %s", requestError.Error())
	}

	response, responseError := _client.Do(request)
	if responseError != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur envoi requête - %s", responseError.Error())
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil,
			response.StatusCode,
			fmt.Errorf("Erreur réponse - %s", response.Status)
	}

	var countries []Country

	decodeError := json.NewDecoder(response.Body).Decode(&countries)
	if decodeError != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur decode données - %s", decodeError.Error())
	}

	return &countries, response.StatusCode, nil
}

func GetCountryByAlpha3(code string) (*Country, int, error) {

	_client := http.Client{Timeout: 5 * time.Second}

	request, requestError := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.apicountries.com/alpha/%s", code), nil)
	if requestError != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur préparation requête - %s", requestError.Error())
	}

	response, responseError := _client.Do(request)
	if responseError != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur envoi requête - %s", responseError.Error())
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil,
			response.StatusCode,
			fmt.Errorf("Erreur réponse - %s", response.Status)
	}

	var country Country

	// body, _ := io.ReadAll(response.Body)
	// fmt.Println(string(body))

	decodeError := json.NewDecoder(response.Body).Decode(&country)
	if decodeError != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur decode données - %s", decodeError.Error())
	}

	return &country, response.StatusCode, nil
}

func GetCountryByName(name string) (*[]Country, int, error) {

	_client := http.Client{Timeout: 5 * time.Second}

	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.apicountries.com/name/%s", name), nil)
	if requestErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur préparation requête - %s", requestErr.Error())
	}

	response, responseErr := _client.Do(request)
	if responseErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur envoi requête - %s", responseErr.Error())
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil,
			response.StatusCode,
			fmt.Errorf("Erreur réponse - %s", response.Status)
	}

	var country []Country

	decodeErr := json.NewDecoder(response.Body).Decode(&country)
	if decodeErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur decode - %s", decodeErr.Error())
	}

	return &country, response.StatusCode, nil

}

func GetCountryByLang(lang string) (*[]Country, int, error) {

	_client := http.Client{Timeout: 5 * time.Second}

	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.apicountries.com/lang/%s", lang), nil)
	if requestErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur préparation requête - %s", requestErr.Error())
	}

	response, responseErr := _client.Do(request)
	if responseErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur envoi requête - %s", responseErr.Error())
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil,
			response.StatusCode,
			fmt.Errorf("Erreur réponse - %s", response.Status)
	}

	var country []Country

	decodeErr := json.NewDecoder(response.Body).Decode(&country)
	if decodeErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur decode - %s", decodeErr.Error())
	}

	return &country, response.StatusCode, nil
}
