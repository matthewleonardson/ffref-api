# Fantasy Football Reference API

Fantasy Football Reference API is a project I've been working on since the summer of 2024. It's an attempt to mimic the [Sports Reference](https://www.sports-reference.com/?utm_source=pfr&utm_medium=sr_xsite&utm_campaign=2023_01_srnav&__hstc=223721476.50912e0c0460e9e7bf8f30fb95a57dc7.1731401264892.1731401264892.1731401264892.1&__hssc=223721476.1.1731401264892&__hsfp=2117850730) style of sports archiving, but for Fantasy Football leagues (or any sort of Fantasy league or even anything Fantasy adjacent!) This repository contains a simplified version of the project, containing few endpoints. I am making this public not for the sake of open sourcing the API my friends and I use, but rather to provide a launchpoint for anyone else working on a similar project.

## Installation 
This API runs as a Docker container. The `Dockerfile` and `docker-compose.yml` are in the repository. Also included is `schema.sql`, which is a .sql file that will create the PostgreSQL schema that this API uses. However, the code should be readable enough to modify and adjust to any existing database you may have.

## Environment Variables
This project expects a `.env` file containing the **DB_USER**, **DB_PASSWORD**, **DB_NAME** and optionally, **API_PORT**.