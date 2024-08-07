# go-url-shorter

## Description
Un URL Shortener est un outil qui permet de transformer une URL longue et complexe en une URL courte et facile à partager. Il fonctionne en générant une clé unique pour l'URL longue et en la stockant dans une base de données. Lorsque l'utilisateur visite l'URL courte, il est redirigé vers l'URL longue d'origine.

## Fonctionnalités
- **Raccourcissement d'URL**: Accepte une URL longue en entrée et génère une URL courte et unique.
- **Base de données**: Stocke les URL longues et courtes dans une base de données Redis.
- **Redirection**: Redirige les utilisateurs vers l'URL longue lorsqu'ils visitent l'URL courte.
- **Statistiques**: Affiche des statistiques basiques comme le nombre de liens raccourcis et les clics sur chaque lien.

# Equipe de developpement
- WELTMANN Jeremy
- NIO Tristan
- MAMACHE Rayan

# Lancement avec docker

```
docker-compose up --build
```