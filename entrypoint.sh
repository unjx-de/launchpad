#!/bin/sh

cat logo.txt
export APP=app

if [ -n "$PUID" ] || [ -n "$PGID" ]; then
  export USER=appuser
  export HOME=/"$APP"

  addgroup -g "$PGID" "$USER" && adduser -u "$PUID" -G "$USER" -h $HOME -D "$USER"
  chown "$USER":"$USER" $HOME -R

  printf "UID: %s GID: %s\n\n" "$PUID" "$PGID"

  su -c - $USER ./"$APP"
else
  printf "WARNING: %s Running docker as root\n" "$(date +'%Y/%m/%d %T')"
  ./"$APP"
fi
