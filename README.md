# simple-pre-commit

## 用法

在`Makefile`文件中使用，比如下面这个例子
```
install:
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  go install github.com/brenner8023/simple-pre-commit@latest
  go mod tidy

lint:
  golangci-lint run --timeout 60s --max-same-issues 50 ./...

pre-commit:
  @make lint
```

## 注意

需要把`${GOPATH}/bin`添加到全局变量中
比如编辑`~/.bashrc`
```bash
export PATH=$PATH:/Users/brenner/go/bin
```
