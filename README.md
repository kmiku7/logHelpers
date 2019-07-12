# logHelpers
logHelpers
日志库的辅助工具。

20190712  
废弃该 repo，当时写这段的代码的目的有两点：

1. 解耦logger与其他代码：`Logger interface`；
2. 在logger输出里增加统一的信息：`prefixLoggerWrapper struct`，用在 api handler 里，每次处理时生成一个唯一ID，添加到日志里，方便查看日志。

对于第一点，单独提供一个 Logger interface 并不是一个太好的实践，由每个使用方在实现自己系统时单独定义一个 interface 即可。  
对于第二点，使用 `github.com/sirupsen/logrus` 库就可以实现；这里的实现也有问题，比如内部嵌套的 Logger 对象会通过 call depth 获取输出日志位置的代码信息并输出，大部分日志库也并没有暴露出 call depth 的配置接口，这里额外封装了一层，导致了输出的位置信息是 `prefixLoggerWrapper` 的代码位置，无法快速直接定位日志的输出位置。

目前来看，日志库的架构形式基本上是前端加后端的实现。前端负责信息格式化，比如添加时间信息，日志输出位置信息等；后端负责代码的输出或持久化。  
在具体项目里，我们可以使用 `github.com/sirupsen/logrus` 作为日志前端进行日志输出的调用，使用 `github.com/kmiku7/golog` 里的 `FileBackend` 作为后端，实现日志持久化，以及日志文件切割等功能。
