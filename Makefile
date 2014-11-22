all:
	mkdir openseadragon
	curl -sL https://github.com/openseadragon/openseadragon/releases/download/v1.1.1/openseadragon-bin-1.1.1.tar.gz | tar zxf - -C openseadragon --strip 1
	gom install
	gom build
	./_vendor/bin/nrsc deepzoom-osd-server openseadragon
