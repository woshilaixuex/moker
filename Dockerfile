FROM golang:1.19
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o user_api ./app/user/cmd/api
RUN go build -o user_rpc ./app/user/cmd/rpc
#   复制配置文件
COPY ./app/user/cmd/api/etc/user-api.yaml etc/user-api.yaml
COPY ./app/user/cmd/rpc/etc/user-rpc.yaml etc/user-rpc.yaml

#   复制启动脚本
COPY ./app/start.sh /app/start.sh
RUN chmod +x /app/start.sh
EXPOSE   12110
#   使用启动脚本启动两个服务
CMD ["/app/start.sh"]
