# Mi Aplicación en Go

Sistema de Administración de Vacunacion

Este sistema de administración de medicamentos es una aplicación Rest escrita en Go que permite gestionar la información de las vacunas disponibles en una farmacia.la aplicaion podra manejar las solicitudes HTTP y el paquete jwt-go para manejar los tokens JWT.

La aplicación incluye dos componentes principales:

- Un servidor HTTP que gestiona las solicitudes HTTP y las respuestas.
- Una base de datos que almacena la información de los medicamentos.

## Instalación

Requisitos
Para instalar y ejecutar la aplicación, necesitarás tener instalado lo siguiente:

Docker Debido a que se requiere  una base de datos, Docker te permite crear un contenedor con Postgres sin preocuparte por la configuración del sistema oper
1. [Go](https://golang.org/dl/) (versión  1.16 o superior)


2. Comandos para la aplicacion
docker-compose up -d
