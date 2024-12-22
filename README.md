# Stage chez OpenSVC, partie oc3

> **Important** Si ce n'est pas déja fait, voir avant le README.md de [om3](https://github.com/Alexandre-Meunier/om3) pour avoir un meilleur contexte.
>
> **_NOTE:_** Ce dépot a été "fork" via le projet principal [oc3](https://github.com/opensvc/oc3) dans le cadre de mon stage de BUT 2 Informatique.


oc3
===

Au cours de mon stage, j'ai pu contribuer au développement d'un autre projet associé à om3, nommé “collecteur”. Il permet de collecter les états, de visualiser et piloter l’ensemble des clusters d’un client.

Un collecteur se compose d’une base de données, d’un cache Redis, d’une API REST et d’une application web en python.

Ce projet oc3 a été lancé en réponse à des problèmes de performance remontés par les utilisateurs de l'ancienne version oc2 qui était encore en Python.



🎯Mes missions
============

Ma mission pour ce projet a été de développer la chaîne de réception des données manipulées dans mon premier développement om3. 

Cette chaîne se compose d’un gestionnaire d’API qui pousse les données reçues dans une file d’attente tenue par Redis, et d’un “worker”, c’est-à-dire un daemon qui dépile les messages dans cette queue pour leur appliquer un traitement long : le dispatch des données en base de données relationnelle.

Pour cette mission, je devais donc d’abord comprendre le fonctionnement de la base clé-valeur Redis et utiliser les acquis de ma première mission pour le développement du gestionnaire d’API oc3.


💻Technologies et pratiques utilisées
======================

- Golang
- API REST
- Daemon Linux
- Redis
- Containers
- Intégration continue (CI)
- IDE Goland (Jetbrains)


📈Statistiques
============

- 10 commits <span style="color: green;">496 ++ </span> <span style="color: red;">164 --</span>
- 2 Pull Requests


