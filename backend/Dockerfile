FROM golang:latest as build
WORKDIR /usr/src/app/llm_api
COPY . .
ENV CGO_ENABLED=0
RUN go mod tidy
RUN go build -o llmapi ./src

FROM gcr.io/distroless/static-debian11
COPY --from=build /usr/src/app/llm_api/llmapi /
EXPOSE 3035
CMD ["/llmapi"]
