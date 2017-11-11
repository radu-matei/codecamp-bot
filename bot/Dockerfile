FROM node:8.2.1-alpine

COPY . /code/bot
WORKDIR /code/bot

RUN npm install
RUN npm install -g typescript@latest
RUN tsc

EXPOSE 3978

CMD ["node", "build/app.js"]
