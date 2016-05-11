
.PHONY: build clean

build: file.pb.go
   
   
file.pb.go: file.proto
	protoc --gogoslick_out=. $^
     
     
clean:
	rm -f file.pb.go