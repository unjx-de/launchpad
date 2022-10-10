# Launchpad

A blazing fast start-page for your services written with Nuxt3 and Go.

![](https://img.shields.io/badge/Framework-Nuxt3-informational?style=for-the-badge&logo=nuxtdotjs&color=4FC08D)
![](https://img.shields.io/badge/Language-Typescript-informational?style=for-the-badge&logo=typescript&color=3178C6)
![](https://img.shields.io/badge/Language-Go-informational?style=for-the-badge&logo=go&color=00ADD8)

[![Build Status](https://build.unjx.de/buildStatus/icon?style=flat-square&job=launchpad%2Fmain)](https://build.unjx.de/job/launchpad/job/main/)

[https://hub.docker.com/r/unjxde/launchpad](https://hub.docker.com/r/unjxde/launchpad)

## Dark Mode

<img style="border-radius:0.5rem" src="https://filedn.eu/lhdsENsife1QUzPddOpRjb5/screenshot_dark_V3.png" alt="dark_mode" width="500"/>

## Light Mode

<img style="border-radius:0.5rem" src="https://filedn.eu/lhdsENsife1QUzPddOpRjb5/screenshot_light_V3.png" alt="light_mode" width="500"/>

## docker-compose example:

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
      - ALLOWED_HOSTS=https://home.example.com
      # available log-levels: trace,debug,info,warn,error,fatal,panic
      - LOG_LEVEL=info
      # location: Stuttgart
      - LOCATION_LATITUDE=48.644929601442485
      - LOCATION_LONGITUDE=9.349618464869025
      # create account here:
      # https://home.openweathermap.org/users/sign_up
      - OPEN_WEATHER_KEY=thisIsNoFunctioningKey
      - OPEN_WEATHER_UNITS=metric
      - OPEN_WEATHER_LANG=en
      # live system graph can be turned off
      - LIVE_SYSTEM=true
    volumes:
      - ./storage:/app/storage
    ports:
      - "127.0.0.1:4000:4000"
```
