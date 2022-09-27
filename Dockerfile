FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go  ./
COPY /pkg/auth/jwt.go   ./pkg/auth/
COPY /pkg/config/app.go ./pkg/config/

COPY /pkg/controllers/newquestion-controller.go ./pkg/controllers/
COPY /pkg/controllers/secure-controller.go ./pkg/controllers/
COPY /pkg/controllers/token-controller.go ./pkg/controllers/
COPY /pkg/controllers/user-controller.go ./pkg/controllers/

COPY /pkg/middlewares/auth.go ./pkg/middlewares/
COPY /pkg/middlewares/uploadfile.go ./pkg/middlewares/

COPY /pkg/models/question.go ./pkg/models/
COPY /pkg/models/score.go ./pkg/models/
COPY /pkg/models/user.go ./pkg/models/



RUN go build -o /test-project main.go

EXPOSE 8080

CMD [ "/test-project" ]
