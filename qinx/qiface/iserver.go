package qiface

type IServer interface {
    Start()

    Stop()

    Serve()
}
