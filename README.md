# awx-go

AWX (Ansible Tower) SDK for the Go programming language.

This SDK has been developed against AWX 14.0.0.

## Installing

```
go get -u github.com/mrcrilly/goawx
```

## Example

We can simply import awx-go and call its services, such as PingService:

```
import (
    "log"
    awxGo "github.com/Colstuwjx/awx-go"
)

func main() {
    awx := awxGo.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := awx.PingService.Ping()
    if err != nil {
        log.Fatalf("Ping awx err: %s", err)
    }

    log.Println("Ping awx: ", result)
}
```

More examples can be found at [here](https://github.com/Colstuwjx/awx-go/tree/master/examples).

## Roadmap

awx-go is still in development, and its roadmap could be found at [here](https://github.com/Colstuwjx/awx-go/blob/master/ROADMAP.md).

## Contribute

There are many ways to contribute to awx-go.

* Submit bugs via [Github issues](https://github.com/Colstuwjx/awx-go/issues);
* Submit a [pull request](https://github.com/Colstuwjx/awx-go/pulls) for fixes or features;
* Mail [me](mailto:wjx_colstu@hotmail.com)
