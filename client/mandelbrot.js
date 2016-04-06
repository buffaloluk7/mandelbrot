var MandelbrotService = function(){
    if (!(this instanceof MandelbrotService )) return new MandelbrotService  ();

    var webSocket = null;

    this.OpenWebSocket = function(initialSpecs, callback){
        webSocket = new WebSocket('ws://localhost:8080/mandelbrot', []);

        webSocket.onerror = function (error) {
            console.log('WebSocket Error ' + error);
        };
        webSocket.onmessage = function (e) {
            callback(e.data);
        };
        webSocket.onopen = function(e){
            var webSocketArguments = generateRequestArgument(initialSpecs);
            console.log(webSocketArguments);
            webSocket.send(webSocketArguments);
        }
         webSocket.onclose = function(e) {
            console.log("DISCONNECTED");
        }
    }

    this.GetMandelbrot = function(specs) {
        if(webSocket == null) return;
        webSocket.send(generateRequestArgument(specs));
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
            cReal - realOffset, cImaginary - imaginaryOffset, cReal + realOffset, cImaginary + imaginaryOffset);
    }


};

var InputViewModel = new function(mandelbrotService, mandelbrotSpecCalculationService){

    getZoomInPercentage = function(){ return 90; }
    getDefaultSpecs = function()
    {
        return{
            width: getMandelbrotView().offsetWidth,
            height: getMandelbrotView().offsetHeight,
            iterations: 100,
            minR: -3,
            minI: -1.5,
            maxR: 1,
            maxI: 1.5
        }
    }
    getMandelbrotView = function(){ return document.getElementById("mandelbrotView");}

    var currentSpecs = getDefaultSpecs();

    var ClickOnMandelbrot = function (event)
    {
        console.log(currentSpecs);

        var e = event.target;
        var dim = e.getBoundingClientRect();
        var x = event.clientX - dim.left;
        var y = event.clientY - dim.top;

        mandelbrotService.GetMandelbrot(currentSpecs);

        zoomedSpecs = mandelbrotSpecCalculationService.Calculate(currentSpecs, x, y, getZoomInPercentage());

        currentSpecs = zoomedSpecs;
    };

    this.Init = function(){

        var defaultSpecs = getDefaultSpecs();
        mandelbrotService.OpenWebSocket(defaultSpecs, function(data){
            getMandelbrotView().src = "data:image/jpg;base64," + data;
        });


        getMandelbrotView().onclick = ClickOnMandelbrot;
    };

}(new MandelbrotService(), new MandelbrotSpecCalculationService ());