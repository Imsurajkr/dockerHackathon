# pull official base image
FROM node:12.18.4-alpine3.11

MAINTAINER pankaj
# set working directory
WORKDIR /app
 
# install app dependencies
COPY package.json ./
COPY package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.0 -g --silent

# add app
COPY . ./
EXPOSE 3000

# start app
CMD ["yarn", "start"]
