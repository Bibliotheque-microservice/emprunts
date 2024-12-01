# Étape 1: Utiliser une image de Go pour le build
FROM golang:1.20-alpine AS builder

# Étape 2: Installer les dépendances de base
RUN apk --no-cache add ca-certificates git

# Étape 3: Créer un répertoire de travail
WORKDIR /app

# Étape 4: Copier les fichiers go.mod et go.sum pour télécharger les dépendances
COPY go.mod go.sum ./

# Étape 5: Télécharger les dépendances (dont github.com/lib/pq)
RUN go mod tidy

# Étape 6: Copier le reste du code
COPY . .

# Étape 7: Construire l'application
RUN go build -o app

# Étape 8: Créer une image plus légère pour l'exécution
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]
