FROM golang:1.21.3-alpine AS build
WORKDIR /src

COPY . .

RUN ls -lah /src

# Download dependencies as a separate step to take advantage of Docker's caching.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage bind mounts to go.sum and go.mod to avoid having to copy them into
# the container.
# RUN --mount=type=cache,target=/go/pkg/mod/ \
#     --mount=type=bind,source=go.sum,target=go.sum \
#     --mount=type=bind,source=go.mod,target=go.mod \
#     go mod download

RUN go mod download -x

# This is the architecture you’re building for, which is passed in by the builder.
# Placing it here allows the previous steps to be cached across architectures.
ARG TARGETARCH

# Build the application.

# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage a bind mount to the current directory to avoid having to copy the
# source code into the container.
# RUN --mount=type=cache,target=/go/pkg/mod/ \

RUN ls -lah /src

RUN go mod tidy

RUN go build -o /bin/server .

CMD ["/bin/server"]