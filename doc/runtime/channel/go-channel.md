# channel实现分析

go有锁数据结构，CSP概念的组成因子之一。

## 最佳实践

分为阻塞式chan与非阻塞式chan； 区别在于有无capacity；没有的话channel只是个同步通信工具。

#### close chan

- 对于已关闭的channel，向里面发送内容会panic，再次关闭会panic；

- close一个channel会唤醒所有等待该channel的g，并使其进入grunnable状态；

- 使用`for range`在channel关闭时会自动退出循环；

- 可以使用len与cap函数获取channel的元素数量与buffer容量；

- 如果channel中没有值（这里特指buffer channel）则取值会返回对应类型的0值；区分方法是`x, ok := <-ch`；

#### nil chan

  `var a chan int`不使用make来初始化的chan是nil channel；关闭一个nil channel会导致panic

#### close principle

  - 一个sender，多个receiver，由sender来关闭chan；
  - 多个sender，多个receiver，借助外部程序来做信号广播，类似于done chan；
  - 如果确定不会有goroutine在通信过程中被阻塞，也可以不关闭chan，等待GC

  > [如何关闭channel](https://github.com/ct-zh/goLearn/blob/master/doc/runtime/channel/closeChan/closeChan.go)

## 相关概念

#### 无锁管道

锁是一种常见的并发控制技术，我们一般会将锁分成乐观锁和悲观锁，无锁（lock-free）队列更准确的描述是使用乐观并发控制的队列。乐观并发控制也叫乐观锁，很多人都会误以为乐观锁是与悲观锁差不多，然而它并不是真正的锁，只是一种并发控制的思想。

乐观并发控制本质上是基于验证的协议，我们使用原子指令 CAS（compare-and-swap 或者 compare-and-set）在多线程中同步数据，无锁队列的实现也依赖这一原子指令。

Channel含有用于保护成员变量的互斥锁,然而锁导致的休眠和唤醒会带来额外的上下文切换，如果临界区过大，加锁解锁导致的额外开销就会成为性能瓶颈。

> 临界资源：一次仅允许一个进程使用的共享资源。多个进程必须互斥地访问。
>
> 临界区：访问临界资源的代码。

社区在2014年提出了无锁的Channel实现方案，该方案将Channel分成三个类型：

- 同步channel - 直接传递数据；
- 异步channel - 基于环形缓存的传统生产者消费者模型；
- `chan struct{}` 类型的异步channel - 不需要实现缓冲区和直接发送的语义。


## 源码分析
> [这个人写得太好了，我就不做重复工作了😊](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-channel/#642-%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84)

### 重点标记

- buf chan是通过循环队列来实现的，通过读指针与写指针来获取数据；
- 缓冲区空间不足时g会挂在`sendq`或者`recvq`两个队列下，此时g会被打包成`sudog`
- sudog - 当 g 遇到阻塞，或需要等待的场景时，会被打包成 sudog 这样一个结构。一个 g 可能被打包为多个 sudog 分别挂在不同的等待队列上;

## Reference

- [Go语言Channel的实现原理](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-channel/#fn:6)
- [临界区](https://baike.baidu.com/item/%E4%B8%B4%E7%95%8C%E5%8C%BA/8942134?fr=aladdin)
- 