FROM --platform=linux/amd64 golang:1.24

WORKDIR /app

ENV GOPRIVATE=github.com/SPVJ/fs-common-lib

ARG GIT_USERNAME
ARG GIT_PERSONAL_ACCESS_TOKEN

RUN git config --global url."https://${GIT_USERNAME}:${GIT_PERSONAL_ACCESS_TOKEN}@github.com/SPVJ/fs-common-lib".insteadOf "https://github.com/SPVJ/fs-common-lib"


COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o /document-service-go

EXPOSE 8080

CMD ["/document-service-go"]