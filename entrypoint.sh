#!/bin/sh

cat logo.txt
APP=app

if [ -n "$PUID" ] || [ -n "$PGID" ]; then
  USER=appuser
  HOME=/"$APP"

  if ! grep -q "$USER" /etc/passwd; then
    # Usage: addgroup [-g GID] [-S] [USER] GROUP
    #
    # Add a group or add a user to a group
    #    -g GID       Group id
    addgroup -g "$PGID" "$USER"

    # Usage: adduser [OPTIONS] USER [GROUP]
    # Create new user, or add USER to GROUP
    #    -h DIR       Home directory
    #    -g GECOS     GECOS field
    #    -G GRP       Group
    #    -D           Don't assign a password
    #    -H           Don't create home directory
    #    -u UID       User id
    adduser -h "$HOME" -g "" -G "$USER" -D -H -u "$PUID" "$USER"
  fi

  chown "$USER":"$USER" "$HOME" -R
  printf "UID: %s GID: %s\n\n" "$PUID" "$PGID"
  exec su -c - $USER ./"$APP"
else
  printf "WARNING: Running docker as root\n\n"
  exec ./"$APP"
fi
