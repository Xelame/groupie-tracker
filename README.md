[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)][contributors-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://git.ytrack.learn.ynov.com/YENNOUHI/groupie-tracker">
    <img src="./static/img/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Groupie-Tracker</h3>

  <p align="center">
    Groupie Trackers consiste à recevoir une API donnée et à manipuler les données qu'elle contient, afin de créer un site Web convivial où vous pouvez afficher les informations sur les groupes à travers plusieurs visualisations de données. Ce projet se concentre également sur la création d'événements/actions et sur leur visualisation.
    <br />
    <br /> 
  </p>
</div>

<!-- GETTING STARTED -->
## Getting Started

### Prérequis

* Golang
  ```sh
  go version
  ```

### Installation

1. Clonez le répertoire
   ```sh
   git clone https://git.ytrack.learn.ynov.com/YENNOUHI/groupie-tracker.git
   ```
_Vous devrez voir le répertoire Groupie-Tracker ainsi que son contenue
2. Installez Golang (si ce n'ai pas déjà le cas)
   ```sh
   sudo apt install golang
   ```


<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1.1 Pour lancer le serveur vous devez écrire de votre terminal:
```sh
go run backend.go
```
1.2 ou :
```sh
go run .
```
2 ctrl + clic sur le lien qui s'affiche dans votre terminal afin de l'ouvrir sur le navigateur de votre choix:
```sh
Server Open In http://localhost:8080/home
```
3 le [lien](http://localhost:8080/home) s'ouvrira dans votre navigateur sur la page d'accueil

_Ensuite vous pouvez naviguez sur le site librement, si le site ne s'affiche pas correctement changez de navgateur.

## Artists page
 
 Sur cette [page](http://localhost:8080/artists) vous trouverez tous les artistes et groupes d'artistes.
 Vous avez à disposition une barre de recherche et un filtre afin de rechercher un ou plusieurs groupe(s) en particulier.
 Une fois que vous aurez trouvez votre bonheur il suffit de cliquer sur la carte et vous aurez acces à toutes les information concernant le groupe.

# Locations page

Sur cette [page](http://localhost:8080/locations) vous trouverez la Map utilisant une API Google
Vous trouvez également une barre de recherche afin de savoir où et quand s'est produit vous artistes préférés.
<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- [ ] Lien de l'[API](https://groupietrackers.herokuapp.com/api)


Regardez notre [JamBoard](https://jamboard.google.com/d/1jU73aVwm4rNw5_GPL22fErTuJzbE6COGu4fxu_CPicY/edit?usp=sharing) pour voir comment nous nous sommes organisé


<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

- [ ] Yassine Ennouhi - yassine.ennouhi@ynov.com
- [ ] Alexandre Rolland - alexandre.rolland85@ynov.com
- [ ] Mathis Cébile - mathis.cebile@ynov.com

Project Link: [https://git.ytrack.learn.ynov.com/YENNOUHI/groupie-tracker.git](https://git.ytrack.learn.ynov.com/YENNOUHI/groupie-tracker.git)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://go.dev/
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/github_username/repo_name/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/github_username/repo_name/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/github_username/repo_name/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png