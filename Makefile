MAKEFILE_DIR= $(dir $(lastword $(MAKEFILE_LIST)))

UNZIPED_DIR=101_ObjectCategories
IMAGE_URL=http://www.vision.caltech.edu/Image_Datasets/Caltech101/$(UNZIPED_DIR).tar.gz

BINARY_NAME=ppm-search-engine
TARGET_ARCH = "amd64"
TARGET_OS = "linux"

build: ensure
	@GOARCH=$(TARGET_ARCH) GOOS=$(TARGET_OS) go build -o $(BINARY_NAME) $(MAKEFILE_DIR)*.go

ensure: *.go
	dep ensure --update

register:
	$(MAKEFILE_DIR)/register_ppm.sh $(BINARY_NAME)

dlimg:
	curl $(IMAGE_URL) -o $(MAKEFILE_DIR)tmp/$(UNZIPED_DIR).tar.gz
	tar -zxvf $(MAKEFILE_DIR)tmp/$(UNZIPED_DIR).tar.gz -C $(MAKEFILE_DIR)tmp
	$(MAKEFILE_DIR)jpg_ppm_converter.sh $(MAKEFILE_DIR)tmp/$(UNZIPED_DIR) $(MAKEFILE_DIR)images
	rm -rf ${MAKEFILE_DIR}tmp

clean:
	rm -rf ${MAKEFILE_DIR}tmp/*
	rm -rf ${MAKEFILE_DIR}data/*
	rm -rf ${MAKEFILE_DIR}images/*
