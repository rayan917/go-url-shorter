FROM node:20 as build
WORKDIR /front-app
COPY package*.json ./
COPY nginx-custom.conf ./
RUN npm install
COPY . .
RUN npm run build


FROM nginx:1.19.2-alpine
COPY --from=build /front-app/dist/short-url-interface/ /usr/share/nginx/html
COPY ./nginx-custom.conf /etc/nginx/conf.d/default.conf