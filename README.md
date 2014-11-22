deepzoom-osd-server
===================

Serve deepzoom images with [openseadragon](https://openseadragon.github.io/) from zip archives.

[vips](http://www.vips.ecs.soton.ac.uk/index.php?title=VIPS) (higher than 7.40.6) can be used to create zipped deepzoom images:  
`$ vips dzsave big-image.jpg deepzoom-big-image.zip`


## usage


### build

 * install `gom` package manager

		$ go get github.com/mattn/gom
		
 * make	
 	
 		$ make
	
	
### run

	$ ./deepzoom-osd-server
	
if `./dzi` directory is missing will be created. copy inside a zipped deepzoom `deepzoom-big-image.zip` and browse at

	http://localhost:8080/view/deepzoom-big-image
	
	
	
## License:

* [Public Domain](LICENSE)
