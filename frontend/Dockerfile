FROM node:alpine3.19 as base

WORKDIR /app

FROM base as deps

COPY package*.json ./

COPY package.json package-lock.json* ./

RUN npm install

ENV NEXT_TELEMETRY_DISABLED 1

FROM deps AS inter

COPY . .

EXPOSE 3000

FROM inter AS prod

RUN npm run build

CMD ["npm", "start"]

FROM inter as dev

CMD npm run dev
