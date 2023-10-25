## ipcnv

Is a simple IP address conversion tool written in Go. It provides functionality for converting IP addresses between IPv4 and their corresponding 32-bit signed/unsigned integer representations. This tool can be helpful for various networking and data manipulation tasks where you need to convert IP addresses to a numerical format or vice versa.

## build

```
go build -o build
```

## usage

```
Usage of ipcnv:
  -i string
        input ip address
  -m int
        0 - ipv4 to int32
        1 - int32 to ipv4 (default -1)
  -o string
        output file (optional)
```
