.PHONY: build run clean dostuff
IMAGE_NAME="samplemgr-build-ubuntu"
IMAGE_VERSION="0.1"

makefile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
this_dir :=  $(dir $(makefile_path))
root_dir := $(abspath $(this_dir)/../..)
build_dir := $(root_dir)/build

mounts := -v $(root_dir):/build
#-v "$(root_dir)/build-alpine":/build/build

build:
	docker build -t oletizi/$(IMAGE_NAME):$(IMAGE_VERSION) -t oletizi/$(IMAGE_NAME):latest .

run:
	docker run --rm  $(mounts) oletizi/$(IMAGE_NAME):$(IMAGE_VERSION)

shell:
	docker run --rm -it $(mounts) oletizi/$(IMAGE_NAME):$(IMAGE_VERSION) /bin/sh

clean:
	rm -rf $(build_dir) || true