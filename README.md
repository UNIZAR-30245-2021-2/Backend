# EINAtic - Backend

TP6 de la asignatura de Arquitectura software del Grado en Ingeniería Informática de la EINA (Universidad de Zaragoza).

Esta primera versión está basada en el proyecto de ejemplo:
* https://dev.to/orlmonteverde/api-rest-con-go-golang-y-postgresql-m0o
* https://github.com/orlmonteverde/go-postgres-microblog

En futuras versiones se realizarán cambios añadiendo nuevas funcionalidades.
## Documentación
En la Wiki del repositorio podéis consultar la documentación de la API, sobre todo de cara al Front-end para saber como realizar las peticiones.
Iré añadiendo más documentación a medida que aumente las funcionalidades.

## ¿Qué hace ahora?
Crea las todas las tablas necesarias para la base de datos. 
Pero únicamente se puede interaccionar con las entidades de usuario y publicación.
Se pueden realizar las peticiones básicas CRUD y devuelve el resultado mediante un JSON.

## Antes de empezar
Es necesario tener instalado PostgreSQL con un base de datos llamada "einatic" cuyo usuario y contraseña de acceso sean "gopher:gopher".
También es necesario tener GoLang.

Se van a emplear las siguientes librerías:
* github.com/go-chi/chi

* github.com/joho/godotenv

* github.com/lib/pq

* golang.org/x/crypto

* github.com/dgrijalva/jwt-go

## Software recomendado
Yo he usado GoLang como IDE porque integra el acceso a la base de datos.
También he empleado 'postman' para realizar las pruebas de peticiones.
