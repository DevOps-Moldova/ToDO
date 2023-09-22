FROM node:16-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm install -g @angular/cli
RUN npm install --legacy-peer-deps
EXPOSE 4200
CMD ["npm", "start"]
