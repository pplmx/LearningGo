# Communication

## Reference

[Reference Here](https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/)

## instruction

```shell
cd communication

go run tcp/server.go 12345
go run tcp/client.go 127.0.0.1:12345

go run udp/server.go 12345
go run udp/client.go 127.0.0.1:12345

go run tcp/concurr/server.go 12345
# the following are both ok
go run tcp/client.go 127.0.0.1:12345
nc 127.0.0.1 12345
```
