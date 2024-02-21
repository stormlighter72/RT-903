#!/bin/sh

# Utilisez la variable d'environnement SECRET_MESSAGE pour récupérer le message depuis le Secret
MESSAGE=$(echo $SECRET_MESSAGE | base64 --decode)

# Créez le fichier avec le message
echo "$MESSAGE" > /shared-data/message.txt

# Sortez pour que l'InitContainer se termine
