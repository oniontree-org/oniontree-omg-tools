include Build.def

.PHONY: all rebuild clean

all:
	go build -v -ldflags="$(LDFLAGS)" -gcflags="$(GCFLAGS)" -asmflags="$(ASMFLAGS)" -o "$(TARGET)"

rebuild:
	go build -a -v -ldflags="$(LDFLAGS)" -gcflags="$(GCFLAGS)" -asmflags="$(ASMFLAGS)" -o "$(TARGET)"

clean:
	$(RM) $(TARGET)

