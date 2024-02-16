FROM node:20

EXPOSE 3000

ADD . /diario
WORKDIR /diario

COPY package.json ./
COPY package-lock.json ./
COPY ./ ./
RUN npm i


CMD ["npm", "run", "start"]