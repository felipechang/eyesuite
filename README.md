# EyeSuite

Image recognition server for NetSuite

## Getting started

1) Create the following docker-compose.yml file

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

2) Run `docker-compose up --detach`

3) Wait a few minutes for the app to build/start

4) Go to http://localhost:8080

## Usage

### Login:

Login using admin:admin

![login](assets/login.png)

### Configuration:

You'll be shown the intro screen

![intro](assets/intro.png)

## Connection:
- Settings > Connection

Connection is where you enter the NetSuite endpoint and token info. 
This is used to communicate with SuiteTalk.

![connection](assets/connection.png)

### Plugins:

- Settings > Plugins

Plugins are custom image readers which you can activate/deactivate on profiles

![plugins](assets/plugins.png)

### Profiles:

- Settings > Profiles

Inside Profiles, you can configure:
- Which plugin to use to process image reading
- Field mappings to add extra information
- Template used to post data to NetSuite, you can add extra field data here inside double brackets `{{}}`

** Template variables starting with an uppercase letter are handled by the server, so they shouldn't be removed ** 

![profiles](assets/profiles.png)

### Users:

- Settings > Users

Create or edit users, and enable access. Passwords are encoded so only the user knows it.

![plugins](assets/users.png)


### Reader:

- Reader
- 
How to use the image reader?

![reader1](assets/reader1.png)

Select a profile from the dropdown

![reader2](assets/reader2.png)

Select an image to process

![reader3](assets/reader3.png)

Click the Read button to process the image

![reader4](assets/reader4.png)

Click the Submit to post image and field data to NetSuite

## License

MPL-2.0