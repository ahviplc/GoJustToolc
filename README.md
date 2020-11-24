# GoJustToolc

-------------------------------------------------------------------------------

## 项目banner

```markdown

   _____            _           _ _______          _      
  / ____|          | |         | |__   __|        | |     
 | |  __  ___      | |_   _ ___| |_ | | ___   ___ | | ___ 
 | | |_ |/ _ \ _   | | | | / __| __|| |/ _ \ / _ \| |/ __|
 | |__| | (_) | |__| | |_| \__ \ |_ | | (_) | (_) | | (__ 
  \_____|\___/ \____/ \__,_|___/\__||_|\___/ \___/|_|\___|                                                                                              
                      Full Of ❤Love❤                                                           
```

banner生成网址:
> http://patorjk.com/software/taag/#p=testall&f=Graffiti&t=GoJustToolc

-------------------------------------------------------------------------------

## fork me
ahviplc/GoJustToolc: ❤GoJustToolc > Go Tools For U (You) ❤
> https://github.com/ahviplc/GoJustToolc

GoJustToolc: ❤GoJustToolc > Go Tools For U (You) ❤
> https://gitee.com/ahviplc/GoJustToolc

-------------------------------------------------------------------------------

## who is who
> 我的Java语言的JustToolc项目地址:

```markdown
ahviplc/JustToolc: ❤JustToolc > Java Tools For U (You) ❤
https://github.com/ahviplc/JustToolc

JustToolc: ❤JustToolc > Java Tools For U (You) ❤
https://gitee.com/ahviplc/JustToolc
```

-------------------------------------------------------------------------------

## slogan
```markdown
❤GoJustToolc > Go Tools For U (You)❤
```

-------------------------------------------------------------------------------

## 如何安装运行测试

### 安装方式 使用go get

```markdown
To start using GoJustToolc, install Go and run go get:
```
> $ go get -u github.com/ahviplc/GoJustToolc

### 使用go mod

> 如果是 Go modules 的项目，您还可以直接编辑 go.mod 文件

> 可使用代理:

> https://mirrors.aliyun.com/goproxy
 
> https://goproxy.cn

> https://goproxy.io

```markdown
module your_project_name

go 1.14

require (
    github.com/ahviplc/GoJustToolc v0.3.0
)
```
> 也可使用 github.com/ahviplc/GoJustToolc latest 获取最新版

### 执行测试main.go
```markdown
使用go mod运行:
1: git clone https://github.com/ahviplc/GoJustToolc.git

2: go mod download

3: go run main.go
```

-------------------------------------------------------------------------------

## 添砖加瓦

### 分支说明

GoJustToolc的源码分为两个分支，功能如下：

| 分支       | 作用                                                          |
|-----------|---------------------------------------------------------------|
| master | 主分支，release版本使用的分支，与中央库提交的jar一致，不接收任何pr或修改 |
| dev    | 开发分支，默认为下个版本的SNAPSHOT版本，接受修改或pr                 |

### 提供bug反馈或建议

提交问题。

- [Gitee issue](https://gitee.com/ahviplc/GoJustToolc/issues)
- [Github issue](https://github.com/ahviplc/GoJustToolc/issues)


### 贡献代码的步骤

1. 在Gitee或者Github上fork项目到自己的repo
2. 把fork过去的项目也就是你的项目clone到你的本地
3. 修改代码（记得一定要修改dev分支）
4. commit后push到自己的库（dev分支）
5. 登录Gitee或Github在你首页可以看到一个 pull request 按钮，点击它，填写一些说明信息，然后提交即可。
6. 等待维护者合并

### PR遵照的原则

需要提交的pr（pull request）符合一些规范，规范如下：

1. 注释完备，尤其每个新增的方法应按照Go文档规范标明方法说明、参数说明、返回值说明等信息，必要时请添加单元测试，如果愿意，也可以加上你的大名。
4. 请pull request到`dev`分支。

-------------------------------------------------------------------------------

## 维护者部署新版本步骤
```markdown
1.切换分支
git checkout master

2.git merge
git merge --no-ff -m "release v0.#.#" dev

3.git push
git push origin master

4.然后再次切换到dev
git checkout dev

5.执行一次git merge master 同步主库最新提交日志,变更到和主库一致的最新git状态
git merge master

6.接下来在分支dev进行愉快的再次开发维护吧.
QAQ
```
-------------------------------------------------------------------------------

## about me
```markdown
By LC
寄语:一人一世界,一树一菩提!~LC
Version 0.1.0 From 20200625
```

