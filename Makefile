MYDIR = .
all: $(MYDIR)/*.go
	for file in $^ ; do \
		go build -ldflags "-s -w" $${file} ; \
	done