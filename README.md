# Songs Indexer

## Requisitos

* Docker
* Docker Compose


## Endpoints

Puedes chequear la documentación generada para Swagger en la URL `/docs/index.html` una vez que el servicio esté corriendo.


>Nota: debes estar autenticado enviando el token JWT suministrado por los siguientes endpoints a través de los headers.
> ```json
> {
>   "Authorization": "Bearer ${token}"
> }
> ```


## Endpoints de autenticación

### Autenticación

### `/auth/register`

Registro de usuarios

#### Método: `POST`

#### Input

```json
{
  "email": "string",
  "password": "string"
}
```


#### Output

```json
{
  "status": "string"
}
```

### `/auth/login`

Obtención del token

#### Método: `POST`

#### Input

```json
{
  "email": "string",
  "password": "string"
}
```


#### Output

```json
{
  "code": "number",
  "expire": "string",
  "token": "string"
}
```

### `/auth/refresh_token`

Método para obtener un nuevo token

#### Método: `GET`

#### Headers

```json
{
  "Authorization": "Bearer ${token}"
}
```

#### Input

**N/A**

#### Output

```json
{
  "code": "number",
  "expire": "string",
  "token": "string"
}
```

## Pasos para correr el servicio

### Producción

Para compilar el proyecto y ejecutar en producción sólo necesitas correr en una terminal:

```bash
# Construye la imagen de la API y el proxy NGINX
docker-compose build
# Levantamiento de todos los servicios
docker-compose up
```

Y el proyecto estará corriendo en el puerto 80

### Desarrollo

Para levantar los servicios necesarios para el desarrollo necesitas correr en la terminal:

```bash
docker-compose up cache db -d
```

Y copia el archivo `.env.example` como `.env` para tener disponible las variables de entorno necesarias en tiempo de ejecución.

Luego simplemente corre:

```bash
go run ./cmd/server
```