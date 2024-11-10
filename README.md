# Mutants API

## Instrucciones para ejecutar localmente

1. Clonar el repositorio
2. Crear un archivo `.env` basado en `.env.example` y agregar las variables de entorno
3. Obtener las dependencias

```bash
go mod tidy
```

4. Para ejecutar localmente, es necesario tener instalado `docker` y correr el siguiente comando para levantar la base de datos

```bash
docker run --name mutants -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=mutants-db -p 5432:5432 -d postgres
```

5. Una vez levantada la base de datos, ejecutar la aplicación

```bash
go run main.go
```

## Correr tests

1. Tener levantada la base de datos
2. Ejecutar el comando `go test -v` en los siguientes archivos:

- `routes`
- `services`

Ejemplo

```bash
cd routes
go test -v
```

## Endpoints

### Base URL

Local: `localhost:8080`

Nube: `https://compulsory-catlaina-san-valen-c2665b5a.koyeb.app`

### Endpoints

- `POST /mutant`
- `GET /stats`

### POST {base_url}/mutant

Recibe un JSON con el siguiente formato

```json
{
  "dna": ["ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"]
}
```

Retorna HTTP 200 si el ADN es mutante y HTTP 403 si no lo es

### GET {base_url}/stats

Retorna un JSON con el siguiente formato

```json
{
  "count_mutant_dna": 40,
  "count_human_dna": 100,
  "ratio": 0.4
}
```
