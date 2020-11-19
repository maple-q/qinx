package main

import (
    "qinx/qnet"
)


func main() {
    s := qnet.NewServer("[qinx v0.1]")
    s.Serve()
}
