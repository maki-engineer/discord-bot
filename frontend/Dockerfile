FROM node:22

WORKDIR /app
COPY package*.json ./

RUN npm install
RUN apt-get update

COPY . .
CMD ["npm", "run", "dev"]
