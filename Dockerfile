FROM golang:1.20
WORKDIR /src/
COPY ./ ./
RUN chmod +x ./build.sh && ./build.sh
