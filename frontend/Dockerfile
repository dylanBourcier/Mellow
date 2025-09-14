## syntax=docker/dockerfile:1.18.0

ARG NV=20
FROM node:${NV}-alpine AS deps
# No spying u baddie
ENV NEXT_TELEMETRY_DISABLED=1
WORKDIR /app
COPY package*.json ./
RUN --mount=type=cache,target=/root/.npm \
	npm ci

FROM node:${NV}-alpine AS builder
ENV NODE_ENV=production \
	NEXT_TELEMETRY_DISABLED=1
WORKDIR /app
COPY --from=deps /app/node_modules /app/node_modules
COPY . .
# Overwrite backend origin at build
ARG BACKEND_ORIGIN=http://localhost:3225
ENV BACKEND_ORIGIN=${BACKEND_ORIGIN}
RUN --mount=type=cache,target=/root/.npm \
	npm run build

# Dev
FROM node:${NV}-alpine AS dev
WORKDIR /app
ENV NEXT_TELEMETRY_DISABLED=1
COPY package*.json ./
RUN --mount=type=cache,target=/root/.npm \
	npm ci
COPY . .
EXPOSE 3000
CMD [ "npm", "run", "dev" ]

# Our headless knight
FROM gcr.io/distroless/nodejs20-debian12:nonroot AS final
ENV NODE_ENV=production \
	NEXT_TELEMETRY_DISABLED=1 \
	PORT=3000
WORKDIR /app
COPY --chown=65532:65532 --from=builder /app/public /app/public
COPY --chown=65532:65532 --from=builder /app/.next/standalone /app
COPY --chown=65532:65532 --from=builder /app/.next/static /app/.next/static
USER 65532:65532
EXPOSE 3000
CMD [ "server.js" ]