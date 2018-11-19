# dockerfile

# base image
FROM node:10.8

# set working directory
RUN mkdir /usr/src/app
WORKDIR /usr/src/app

RUN npm install react-scripts@1.1.0 -g
RUN npm install

# start app
CMD ["npm", "start"]
