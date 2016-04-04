var MandelbrotService = function(inputProvider){
    if (!(this instanceof MandelbrotService )) return new MandelbrotService  ();

    this.GetMandelbrot = function(specs, callback){
        var webSocket = new WebSocket('ws://localhost:8080/mandelbrot', []);
        webSocket.onerror = function (error) {
            console.log('WebSocket Error ' + error);
        };
        webSocket.onmessage = function (e) {
            callback(e.data);
        };
        webSocket.onopen = function(e){
            var webSocketArguments = generateRequestArgument(specs);
            console.log(webSocketArguments);
            webSocket.send(webSocketArguments);
        }
    }

    generateRequestArgument = function(specs){
        var nl = function(value) { return value + ";" };
        return nl(specs.width) + nl(specs.height) + nl(specs.minR) + nl(specs.minI) + nl(specs.maxR) + nl(specs.maxI) + nl(specs.iterations);
    }
};

var specFactory = function (width, height, iterations, minR, minI, maxR, maxI) {
        return {
            width: width,
            height: height,
            iterations: iterations,
            minR: minR,
            minI: minI,
            maxR: maxR,
            maxI: maxI
        }
    }

var MandelbrotSpecCalculationService = function(){
    if (!(this instanceof MandelbrotSpecCalculationService)) return new MandelbrotSpecCalculationService();

    this.Calculate = function(oldSpecs, x, y, percentage){

        realRange = oldSpecs.maxR - oldSpecs.minR;
        cReal = parseFloat(x) * (realRange / parseFloat(oldSpecs.width)) + oldSpecs.minR;
        imaginaryRange = oldSpecs.maxI - oldSpecs.minI;
        cImaginary = parseFloat(y) * (imaginaryRange / parseFloat(oldSpecs.height)) + oldSpecs.minI;

        percentage = (percentage > 1) ? percentage / 100 : percentage;

        var realOffset = realRange * percentage / 2;
        var imaginaryOffset = imaginaryRange * percentage / 2;

        var newIteration = ((1 - percentage) + 1)*oldSpecs.iterations;

        return specFactory(oldSpecs.width, oldSpecs.height, newIteration,
            oldSpecs.minR + realOffset, oldSpecs.minI + imaginaryOffset, oldSpecs.maxR - realOffset, oldSpecs.maxI - imaginaryOffset);
    }


};

var InputViewModel = new function(mandelbrotService, mandelbrotSpecCalculationService){

    getZoomInPercentage = function(){ return 90; }
    getDefaultSpecs = function()
    {
        return{
            width: 1000,
            height: 750,
            iterations: 100,
            minR: -3,
            minI: -1.5,
            maxR: 1,
            maxI: 1.5
        }
    }

    var currentSpecs = getDefaultSpecs();
    var mandelbrotContainer = document.getElementById("mandelbrot");

    var ClickOnMandelbrot = function (event)
    {
        console.log(currentSpecs);

        mandelbrotService.GetMandelbrot(currentSpecs, function(imageData){
            mandelbrotContainer.src  = "data:image/jpg;base64," + imageData;
        });

        zoomedSpecs = mandelbrotSpecCalculationService.Calculate(currentSpecs, event.screenX, event.screenY, getZoomInPercentage());

        currentSpecs = zoomedSpecs;
    };

    this.Init = function(){
        mandelbrotContainer = mandelbrotContainer;
        mandelbrotContainer.onclick = ClickOnMandelbrot;
    };

}(new MandelbrotService(), new MandelbrotSpecCalculationService ());