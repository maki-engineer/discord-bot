FROM node:22

WORKDIR /usr/src/app
COPY package*.json ./

RUN npm install
RUN apt-get update
RUN apt-get -y install ffmpeg

COPY . .
CMD ["sh", "-c", "npm run migrate:development && npm run seed:development && npm run start"]
