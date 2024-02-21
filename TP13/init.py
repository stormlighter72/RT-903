import os

file_path = '/shared-data/message.txt'

try:
    with open(file_path, 'r') as file:
        message = file.read().strip()

    if not message:
        raise Exception("Le fichier est vide.")
except Exception as e:
    print(f"Erreur lors de la lecture du fichier : {e}")
    message = os.environ.get('MESSAGE', 'Message par défaut si tout échoue.')

print(f"Message : {message}")
