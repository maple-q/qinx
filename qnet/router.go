package qnet

import "qinx/qiface"

// 实现router时，先嵌入这个基类，再根据需要对这个基类方法进行重写
type BaseRouter struct{}


func (br *BaseRouter) PreHandle(req qiface.IRequest){}
func (br *BaseRouter) Handle(req qiface.IRequest){}
func (br *BaseRouter) PostHandle(req qiface.IRequest){}
