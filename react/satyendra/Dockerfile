# pull official base image
FROM node:13.12.0-alpine as react-img

# set working directory
WORKDIR /var/www/

# Copying package.json and lock file
COPY ./nginx.conf ./
COPY ./package* ./
ENV PATH /node_modules/.bin:$PATH

# Install npm packages
RUN npm install -y --silent

#
# Building application
#
FROM react-img as build

# add `/app/node_modules/.bin` to $PATH
COPY ./src/ ./src/
COPY ./public/ ./public/
RUN npm run build

#
# Testing application
#
# FROM build as test

# Test
# RUN npm run test


#
# Deploy build over HTTP server
#
FROM nginx:1.17.8-alpine

COPY --from=build /var/www/build /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf
COPY /nginx.conf /etc/nginx/conf.d

WORKDIR /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]