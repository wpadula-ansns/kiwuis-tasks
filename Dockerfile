# Etap deweloperski: kompilator Go i kod źródłowy
FROM golang:1.27 AS dev
WORKDIR /src
# Program bez bibliotek systemowych, uruchomi się na każdym obrazie bazowym
ENV CGO_ENABLED=0
RUN go build -o /out/app .
CMD ["go", "run", "."]

# Etap końcowy: tylko skompilowany program
FROM ubuntu:26.04
COPY --from=dev /app /app
EXPOSE 8080
CMD ["/app"]
