go-winres make
go build -trimpath -ldflags "-H=windowsgui -w -s" -tags=noconsole
go build -trimpath -ldflags "-w -s" -o jtrclient-launcher.console.exe
