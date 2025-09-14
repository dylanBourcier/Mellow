# syntax=docker/dockerfile:1.18.0

#  BUild context
ARG TOS=linux
ARG TAC=amd64

FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache build-base
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . ./
# Static binary (GCO nearly cost my life on windows. SO YEAH i switched to my FEDORA VM)
ENV CGO_ENABLED=1
RUN --mount=type=cache,target=/root/.cache/go-build \
	GOOS=${TOS} GOARCH=${TAC} \
	go build -ldflags '-s -w -linkmode external -extldflags "-static"' -o /app/mellow ./
RUN mkdir -p /app/database/migration /app/uploads /app/data

# Dev targer (yeah cuz i'll use docker compose profil, we're pros here)
FROM golang:1.24-alpine AS dev
WORKDIR /app
RUN apk add --no-cache build-base
COPY . .
ENV PORT=3225 \
	DB_PATH=/app/data/social.db \
	MIGRATIONS_PATH=/app/database/migration/sqlite
EXPOSE 3225
CMD [ "sh", "-c", "mkdir -p /app/data /app/uploads && go run ./" ]

# Headless final image (flex on image size), using GCR
# Serious reason : headless distro serve as security measure, in fact they dont have any distro attached
#	so they can't run any commands, can't even exec on it. BUT as we are bound to sqlite we must set a root
#	user so we can actually write on DB, otherwise it will only have readonly perms.
#	So kids stopp using it, so better options exist, at least if you wanna host online.
FROM gcr.io/distroless/static-debian12 AS final
WORKDIR /app
# Here chown 0:0 actually serves so we have full perms on DB, better solutions exists but will require more 
# workload and complexity
COPY --chown=0:0 --from=builder /app/mellow /app/mellow
COPY --chown=0:0 --from=builder /app/database/migration /app/database/migration
COPY --chown=0:0 --from=builder /app/uploads /app/uploads
COPY --chown=0:0 --from=builder /app/data /app/data
# Here declaring our ROOT user for runtime
USER 0:0
ENV PORT=3225 \
	DB_PATH=/app/data/social.db \
	MIGRATIONS_PATH=/app/database/migration/sqlite
EXPOSE 3225
ENTRYPOINT [ "/app/mellow" ]