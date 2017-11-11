FROM radumatei/k8s-client-builder:latest as builder

WORKDIR /go/src/github.com/radu-matei/kube-bot-client

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kube-bot-client .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/radu-matei/kube-bot-client/kube-bot-client .

EXPOSE 80

CMD ["./kube-bot-client"] 