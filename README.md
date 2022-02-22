# EyeSuite

Image recognition server for NetSuite

## Getting started

1) create the following docker-compose.yml file

```
version: '3'
services:

  # Redis DB
  db:
    container_name: eyesuite_redis
    image: redis:latest
    volumes:
      - data:/data
    restart: always

  # OCR server
  server:
    container_name: eyesuite
    image: hardcake/eyesuite:latest
    ports:
      - "8080:8080"
    depends_on:
      - db
    restart: always

volumes:
  data:
```

## Usage

Go start page to http://localhost:8080

## License
MPL-2.0