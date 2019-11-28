FROM node:10.9-slim

RUN apt-get -y update && apt-get install -y git

# définit le dossier 'app' comme dossier de travail
ENV PATH /usr/src/app/node_modules/.bin:$PATH
RUN mkdir /app
RUN chmod +rw /app
WORKDIR /app

COPY view/editor/package.json ./package.json


RUN echo "#!/bin/bash" >> /script.sh
RUN echo "npm install" >> /script.sh
RUN echo "echo \"VUE_APP_SERVER_URL=\$API_HOST\" >> /app/.env" >> /script.sh
RUN echo "echo \"VUE_APP_SERVER_PORT=\$API_PORT\" >> /app/.env" >> /script.sh
RUN echo "npm run serve" >> /script.sh
RUN chmod +x /script.sh

EXPOSE 8080

ENTRYPOINT [ "/script.sh" ]
