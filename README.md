## 协程（Goroutine）
### 进程（Process）
> 指计算机中已运行的程序。进程曾经是分时系统的基本运作单位。在面向进程设计的系统（如早期的UNIX，Linux 2.4及更早的版本）中，进程是程序的基本执行实体；在面向线程设计的系统（如当代多数操作系统、Linux 2.6及更新的版本）中，**进程本身不是基本运行单位，而是线程的容器。**程序本身只是指令、数据及其组织形式的描述，**进程才是程序（那些指令和数据）的真正运行实例**。若干进程有可能与同一个程序相关系，且每个进程皆可以同步（循序）或异步（平行）的方式独立运行。现代计算机系统可在同一段时间内以进程的形式将多个程序加载到存储器中，并借由时间共享（或称时分复用），以在一个处理器上表现出同时（平行性）运行的感觉。同样的，使用多线程技术（多线程即每一个线程都代表一个进程内的一个独立执行上下文）的操作系统或计算机体系结构，同样程序的平行线程，可在多CPU主机或网络上真正同时运行（在不同的CPU上）。
### 线程（Thread）
>
### 协程（Goroutine）
> 
## 参考文档
(golang-goroutine-vs-thread)[https://www.geeksforgeeks.org/golang-goroutine-vs-thread/]
