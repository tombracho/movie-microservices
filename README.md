Before launching the application, you need to have Docker and [Modd](https://github.com/cortesi/modd) (optional) installed on your machine.


## Launching Hashicorp Consul
To run the Hashicorp Consul container, use the following command:

```shell
docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
```

## Launching the application using modd
Use the following command from the root directory of the project
```shell
modd
```
