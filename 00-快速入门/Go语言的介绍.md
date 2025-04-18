## Go语言的介绍
Go 语言是由 Google 开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。
Go 富有表现力、简洁、干净且高效。它的并发机制使编写程序变得容易，从而充分利用多核和联网机器，而其新颖的类型系统则支持灵活和模块化的程序构建。
Go 可以快速编译为机器代码，同时还具有垃圾收集的便利性和运行时反射的强大功能。它是一种快速、静态类型的编译语言，感觉就像是一种动态类型的解释语言。
### 历史
Go 语言起源 2007 年，并于 2009 年正式对外发布。它从 2009 年 9 月 21 日开始作为谷歌公司 20% 兼职项目，即相关员工利用 20% 的空余时间来参与 Go 语言的研发工作。
* 2007 年 9 月 20 日的下午在谷歌山景垅总部的一间办公室里，谷歌的大佬级程序员 RobPike 在等待一个 C++ 项目构建的过程中和谷歌的另外两个大佬级程序员 RobertGriesemer 和 KenThompson 进行了一次有关设计一门新编程语言的讨论。
* 计算机硬件技术更新频繁，性能提高很快。目前主流的编程语言发展明显落后于硬件，不能合理利用多核多 CPU 的优势提升软件系统性能。
* 软件系统复杂度越来越高，维护成本越来越高，目前缺乏一个足够简洁高效的编程语言。
* 企业运行维护很多 C/C++ 的项目，C/C++程序运行速度虽然很快(因为采用静态编译)，但是编译速度却很慢，同时还存在内存泄露的一系列困扰需要解决。
### 里程碑
* 2007年9月20日，由谷歌工程师 Rob Pike，Robert Griesemer 和 Ken Thompson 共同设计开发的一门全新的语言，这是Go的雏型。
* 2009年11月10日，Goolge 将 Go 语言开源正式向全世界发布。
* 2015年8月19日，Go语言1.5正式发布，本次更新移除了 “最后残余的C代码” , 请内存管理方面的权威专家Rick Hudson对GC进行了重新设计(重要修改)。
* 2016年2月17日，Go语言1.6正式发布。
* 2017年2月16日，Go语言1.8正式发布。
* 2018年8月24日，Go语言1.9正式发布。
* 2018年2月16日，Go语言1.10正式发布。
* 2018年8月24日，Go语言1.11正式发布。
* 2019年2月25日，Go语言1.12正式发布。
* 2019年9月03日，Go语言1.13正式发布。
* 2020年2月25日，Go语言1.14正式发布。
* 2020年8月11日，Go语言1.15正式发布。
* 2021年2月16日，Go语言1.16正式发布。
* 2021年8月16日，Go语言1.17正式发布。
* 2022年3月15日，Go语言1.18正式发布。
* 2022年8月02日，Go语言1.19正式发布。
* 2023年2月01日，Go语言1.20正式发布。
* 2023年8月08日，Go语言1.21正式发布。
* 2024年2月06日，Go语言1.22正式发布。
* 2024年8月13日，Go语言1.23正式发布。
* 2025年2月11日，Go语言1.24正式发布。
### 核心成员
![image-1_01](./images/01-介绍/1_01.png)

**Ken Thompson(肯·汤普森)**: 1983年图灵奖获得者，1998年美国国家技术奖得主，他与 Dennis Ritchie 是 Unix 的创始人。同时也是 C语言的主要创始人。
**Rob Pike(罗伯·派克)**: 他是加拿大人，曾经是贝尔实验室的Unix团队和Plan 9操作系统的领导者。他和 Ken Thompson 共事多年，并共创出广泛使用的 UTF-8 编码。(ps: Go语言的图标 gopher 囊地鼠，是他的妻子制作的)
**Robert Griesemer(罗伯特·格瑞史莫)**: 曾协助制作 Java 的 HotSpot 虚拟机和 Chrome 浏览器的 JavaScript V8 引擎的设计者之一。
### Go语言的吉祥物
Go语言的吉祥物是一个囊地鼠，它是由才华横溢的插画家 Renee French 设计的，她也是 Go 设计者之一 Rob Pike 的妻子。
<p><img src="./images/01-介绍/1_02.png" width="30%" clign="center"></p>
### 为什么要学习Go语言？
1. 简单好记的关键词和语法。轻松上手，简单易学。
2. 更高的效率。比Java，C++等拥有更高的编译速度，同时运行效率媲美C，同时开发效率非常高。
3. 生态强大，网络上库很丰富，很多功能使用Go开发非常简单。
4. 语法检查严格，高安全性。
5. 严格的依赖管理，go mod命令。
6. Go拥有强大的编译检查、严格的编码规范和完整的软件生命周期工具，具有很强的稳定性，稳定压倒一切。
7. 跨平台交叉编译，windows就可以编译出mac，linux上可执行的程序。
8. 异步编程复杂度低，易维护，GO 语言中 Channel 设计，异步程序写起来非常自然。
9. 语言层面支持并发，**go关键字（协程）**使得go的并发效率极高。
10. 严格的语法规范，所有程序员写出来的代码都是一样的，对大团队来说，非常友好。
11. Go 的并发、性能、安全性、易于部署等特性，使它很容易成为“云原生语言”。容器和云的使用上非常广。

Go语言设计之初，确认了三大原则: **简洁性，可读性，功能性**。