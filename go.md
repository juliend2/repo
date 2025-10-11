# GO

## initialize a repo

```bash
MODNAME="something"
mkdir $MODNAME && cd $MODNAME
go mod init desrosiers.org/$
```

## compile a file

```bash
go build -o main main.go
```