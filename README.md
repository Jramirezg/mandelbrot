# Mandelbrot
mandelbrot drawer in console

### Overview
The provided Go program generates and visualizes the Mandelbrot set, a famous fractal, in the console. It offers flexibility by allowing users to choose between using dots or characters to represent different levels of iterations in the set. The visualization is configurable through a YAML file and command-line flags, providing control over the image resolution, maximum iterations, and the choice of characters for representation.

### Configuration
The program's configuration is specified in the config.yaml file. You can adjust the following parameters:

* width: Width of the console visualization.
* height: Height of the console visualization.
* maxIter: Maximum number of iterations for calculating the Mandelbrot set.
* contrast: Contrast value for character representation.
* useCharacters: Whether to use characters or dots for visualization.

### Usage
The program generates a console-based visualization of the Mandelbrot set. The choice between using characters or dots for representation can be controlled via the -useCharacters command-line flag. By default, dots are used.

### Examples
Default (dots):

```bash
go run main.go
```
Character representation:

```bash
go run main.go -useCharacters
```
## Notes

.: Represents points in the Mandelbrot set where the escape time (number of iterations) is equal to config.MaxIter. These are points that are considered part of the set.

(space): Represents points outside the Mandelbrot set, i.e., points that diverge to infinity within the specified maximum number of iterations.

:-=+*%@#MW: Represents points that take different numbers of iterations to escape. Each character corresponds to a different range of iterations, with M and W typically representing the points that take the maximum number of iterations (config.MaxIter).