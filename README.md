# Gin backend 

Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API, but with performance up to 40 times faster than Martini. If you need smashing performance, get yourself some Gin.

## ⚡️ Quick start

1. Create a new project with Gin:

```bash
mkdir [folder-name] && cd [folder-name]
go mod init [folder-name]
go get -u github.com/gin-gonic/gin
```

2. Rename `.env.example` to `.env` and fill it with your environment values.

3. Go to API Docs page (Swagger): [127.0.0.1:5000/swagger/index.html](http://127.0.0.1:5000/swagger/index.html)

![Screenshot](https://user-images.githubusercontent.com/11155743/112715187-07dab100-8ef0-11eb-97ea-68d34f2178f6.png)

## 🗄 Template structure


### ./docs

**Folder with API Documentation**. This directory contains config files for auto-generated API Docs by Swagger.


## ⚙️ Configuration

```ini
# .env

# Stage status to start server:
#   - "dev", for start server without graceful shutdown
#   - "prod", for start server with graceful shutdown

# Database settings:
DB_HOST=
DB_DRIVER=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=

# Token
TOKEN_EXPIRED_IN=
TOKEN_MAXAGE=
TOKEN_SECRET=
```
