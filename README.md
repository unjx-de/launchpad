# Launchpad

A blazing fast start-page for your services written with Nuxt3 and Go.

![](https://img.shields.io/badge/Framework-Nuxt3-informational?style=for-the-badge&logo=nuxtdotjs&color=4FC08D)
![](https://img.shields.io/badge/Language-Typescript-informational?style=for-the-badge&logo=typescript&color=3178C6)
![](https://img.shields.io/badge/Language-Go-informational?style=for-the-badge&logo=go&color=00ADD8)

[![Build Status](https://build.unjx.de/buildStatus/icon?style=flat-square&job=launchpad%2Fmain)](https://build.unjx.de/job/launchpad/job/main/)

[https://hub.docker.com/r/unjxde/launchpad](https://hub.docker.com/r/unjxde/launchpad)

## Dark Mode

<img style="border-radius:0.5rem" src="https://filedn.eu/lhdsENsife1QUzPddOpRjb5/Launchpad/screenshot_dark.png" alt="dark_mode" width="500"/>

## Light Mode

<img style="border-radius:0.5rem" src="https://filedn.eu/lhdsENsife1QUzPddOpRjb5/Launchpad/screenshot_light.png" alt="light_mode" width="500"/>

## Lighthouse speed test

<img style="border-radius:0.5rem" src="https://filedn.eu/lhdsENsife1QUzPddOpRjb5/Launchpad/lighthouse.jpeg" alt="lighthouse" width="250"/>

<img style="border-radius:0.5rem" src="https://filedn.eu/lhdsENsife1QUzPddOpRjb5/Launchpad/lighthouse_detail.jpeg" alt="lighthouse" width="500"/>

## How to use

Use the docker compose to spin up the service. The Weather is fetched over a [Current Weather Api Call](https://openweathermap.org/current) with environment variables for the needed parameters. Please refer to the available options as shown in the docker-compose example.

### Example of the bookmarks.json

All Bookmarks are read from a file called `bookmarks.json` located inside the `./storage` folder.
The application will create a default file at startup and will automatically look for changes inside the file.
Changes are printed in stdout when running with `LOG_LEVEL=trace`.

You can specify an icon of a bookmark either by using a link or by using the name of the file located inside the `./storage/icons` folder that is mounted via the docker compose file.
The name and related link can be provided as well.

```json
[
  {
    "NAME": "Github",
    "ICON": "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png",
    "URL": "https://github.com"
  },
  {
    "NAME": "Jenkins",
    "ICON": "jenkins.webp",
    "URL": "https://www.jenkins.io/"
  }
]
```

### Available environment variables with default values

```toml
PORT = 4000
ALLOWED_HOSTS = "http://localhost:4000"
SWAGGER = false

LOG_LEVEL = "info"

LOCATION_LATITUDE = 48.780331609463815
LOCATION_LONGITUDE = 9.177968320179422
OPEN_WEATHER_KEY = ""
OPEN_WEATHER_UNITS = "metric"
OPEN_WEATHER_LANG = "en"

LIVE_SYSTEM = true
```

## A docker-compose example:

```yaml
version: '3.9'

services:

  launchpad:
    image: unjxde/launchpad:latest
    container_name: launchpad
    restart: unless-stopped
    environment:
      # https://docs.linuxserver.io/general/understanding-puid-and-pgid
      - PUID=1000
      - PGID=1000
      # can be multiple hosts, comma separated, no spaces
      - ALLOWED_HOSTS=https://home.example.com
      # available log-levels: trace,debug,info,warn,error,fatal,panic
      - LOG_LEVEL=info
      # location: Stuttgart
      - LOCATION_LATITUDE=48.644929601442485
      - LOCATION_LONGITUDE=9.349618464869025
      # create account here to get free key:
      # https://home.openweathermap.org/users/sign_up
      - OPEN_WEATHER_KEY=thisIsNoFunctioningKey
      # standard, metric or imperial
      - OPEN_WEATHER_UNITS=metric
      # https://openweathermap.org/current#multi
      - OPEN_WEATHER_LANG=en
      # live system graph can be turned off
      - LIVE_SYSTEM=true
    volumes:
      # to mount the bookmarks.json and the icons folder on the system
      - ./storage:/app/storage
    ports:
      - "127.0.0.1:4000:4000"
```
