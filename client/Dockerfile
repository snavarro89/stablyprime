
FROM node:latest AS build-angular
WORKDIR /usr/local/app
COPY ./ /usr/local/app
RUN npm install
RUN node --max_old_space_size=6144 ./node_modules/@angular/cli/bin/ng build --configuration docker --aot

FROM nginx:latest
COPY --from=build-angular /usr/local/app/dist/stablyprime /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY ./nginx.conf /etc/nginx/conf.d
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]

