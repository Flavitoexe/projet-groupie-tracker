# Projet Groupie Tracker - GeoFacts

GeoFacts est un site web permettant de rechercher des pays et obtenir certaines informations sur ces derniers, telles que : 

 - Nom 
 - AlphaCode 2 et 3
 - Capitale 
 - Continent
 - Région du monde
 - Population
 - Superficie
 - Langue(s)
 - Monnaie(s)
 - Domaine de premier niveau
 - Pays frontaliers


## Fonctionnalités
La recherche des pays peut se faire de plusieurs manières : 

 - Par le continent
 - Par le nom
 - Par la/les langue(s) parlée(s)


## Technologies utilisées

 - HTML
 - CSS
 - Golang
 - Api Countries ([Documentation](https://www.apicountries.com/docs))


## Structure du projet

```
geofacts
│   README.md
│
├───assets
│   └───css
│           country_choice.css
│           country_details.css
│           country_lang.css
│           region_choice_style.css
│           style.css
│
├───src
│   │   go.mod
│   │   main.go
│   │
│   ├───controllers
│   │       geofacts.controllers.go
│   │
│   ├───helper
│   │       templates.helper.go
│   │
│   ├───routes
│   │       geofacts.routes.go
│   │       main.routes.go
│   │
│   └───services
│           geofacts.service.go
│
└───templates
        country_choice.html
        country_details.html
        country_lang.html
        error.html
        menu.html
        region_choice.html
```

## Installation et utilisation

1. Cloner le projet 
```
git clone https://github.com/Flavitoexe/projet-groupie-tracker.git
```

2. Ouvrir le dossier pour accéder au main.go
```
cd src
```

3. Lancer le serveur
```
go run main.go
```

4. Accéder à localhost:8080 dans votre navigateur


## Auteur

Flavio GRILLI  
Projet académique réalisé dans le cadre d’études supérieures.