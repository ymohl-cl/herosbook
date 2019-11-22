FROM node:lts-alpine

ARG API_HOST_ARG
ARG API_PORT_ARG

ENV API_HOST=$API_HOST_ARG
ENV API_PORT=$API_PORT_ARG

# installe un simple serveur http pour servir un contenu statique
RUN npm install -g http-server

# définit le dossier 'app' comme dossier de travail
WORKDIR /app

# copie 'package.json' et 'package-lock.json' (si disponible)
COPY package.json package.json

# installe les dépendances du projet
RUN npm install

# copie les fichiers et dossiers du projet dans le dossier de travail (par exemple : le dossier 'app')
COPY .eslintrc.js .eslintrc.js
COPY .browserslistrc .browserslistrc
COPY public public
COPY src src
COPY tsconfig.json tsconfig.json
COPY postcss.config.js postcss.config.js
COPY babel.config.js babel.config.js
RUN echo "VUE_APP_SERVER_URL=${API_HOST}" >> .env
RUN echo "VUE_APP_SERVER_PORT=${API_PORT}" >> .env

# construit l'app pour la production en la minifiant
RUN npm run build

EXPOSE 8080
CMD [ "http-server", "dist" ]