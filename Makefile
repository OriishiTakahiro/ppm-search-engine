MAKEFILE_DIR= $(dir $(lastword $(MAKEFILE_LIST)))

UNZIPED_DIR=101_ObjectCategories
IMAGE_URL=http://www.vision.caltech.edu/Image_Datasets/Caltech101/$(UNZIPED_DIR).tar.gz

dlimg:
	curl $(IMAGE_URL) -o $(MAKEFILE_DIR)tmp/$(UNZIPED_DIR).tar.gz
	tar -zxvf $(MAKEFILE_DIR)tmp/$(UNZIPED_DIR).tar.gz -C $(MAKEFILE_DIR)tmp
	$(MAKEFILE_DIR)jpg_ppm_converter.sh $(MAKEFILE_DIR)tmp/$(UNZIPED_DIR) $(MAKEFILE_DIR)images
	rm -rf ${MAKEFILE_DIR}tmp

clean:
	rm -rf ${MAKEFILE_DIR}tmp/*
	rm -rf ${MAKEFILE_DIR}data/*
	rm -rf ${MAKEFILE_DIR}images/*
