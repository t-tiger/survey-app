FROM node:15.5.0-alpine3.10 AS base
RUN apk add --no-cache yarn
WORKDIR /base
COPY package.json ./
COPY yarn.lock ./
RUN yarn install
COPY . .

FROM base AS build
ENV NODE_ENV=production
WORKDIR /build
COPY --from=base /base ./
RUN yarn build

FROM node:15.5.0-alpine3.10 AS app
ENV NODE_ENV=production
WORKDIR /app
COPY --from=build /build/package.json ./
COPY --from=build /build/yarn.lock ./
COPY --from=build /build/.next ./.next
RUN npm install next

CMD ["npm", "run", "start"]