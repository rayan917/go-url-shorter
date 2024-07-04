# Utilisez une image de base appropriée pour votre application Go
FROM golang:latest

# Définissez le répertoire de travail dans le conteneur
WORKDIR /app

# Copiez le fichier main.go dans le conteneur
COPY . .

# Installez les dépendances
RUN go get -t url-short

# Exécutez les tests
RUN go test -v

RUN go build -o bin .

RUN chmod a+x /bin

# Exposez le port sur lequel votre application écoute
EXPOSE 8080

# Démarrez votre application lorsque le conteneur est lancé
ENTRYPOINT [ "/app/bin" ] 