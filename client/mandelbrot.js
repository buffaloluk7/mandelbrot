var mandelbrotService = function(){

    var webSocket = null;

    var openWebSocket = function(initialSpecs, callback){
        webSocket = new WebSocket('ws://localhost:8080/mandelbrot', []);

        webSocket.onerror = function (error) {
            console.log('WebSocket Error ' + error);
        };
        webSocket.onmessage = function (e) {
            callback(e.data);
        };
        webSocket.onopen = function(e){
            var webSocketArguments = _generateRequestArgument(initialSpecs);
            console.log(webSocketArguments);
            webSocket.send(webSocketArguments);
        }
         webSocket.onclose = function(e) {
            console.log("DISCONNECTED");
        }
    }

    var getMandelbrot = function(specs) {
        if(webSocket == null) return;
        webSocket.send(_generateRequestArgument(specs));
    }

    var _generateRequestArgument = function(specs){
        var nl = function(value) { return value + ";" };
        return nl(specs.width) + nl(specs.height) + nl(specs.minR) + nl(specs.minI) + nl(specs.maxR) + nl(specs.maxI) + nl(specs.iterations);
    };

    return {
        getMandelbrot: getMandelbrot,
        openWebSocket: openWebSocket
    };
}();

var specFactory = function(){
    var create = function (width, height, iterations, minR, minI, maxR, maxI) {
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

    return {
        create: create
    }
}();

var mandelbrotSpecCalculationService = function(specFactory){

    var calculate = function(oldSpecs, x, y, percentage){

        realRange = oldSpecs.maxR - oldSpecs.minR;
        cReal = parseFloat(x) * (realRange / parseFloat(oldSpecs.width)) + oldSpecs.minR;
        imaginaryRange = oldSpecs.maxI - oldSpecs.minI;
        cImaginary = parseFloat(y) * (imaginaryRange / parseFloat(oldSpecs.height)) + oldSpecs.minI;

        percentage = (percentage > 1) ? percentage / 100 : percentage;

        var realOffset = realRange * percentage / 2;
        var imaginaryOffset = imaginaryRange * percentage / 2;

        var newIteration = ((1 - percentage) + 1)*oldSpecs.iterations;

        return specFactory.create(oldSpecs.width, oldSpecs.height, newIteration,
            cReal - realOffset, cImaginary - imaginaryOffset, cReal + realOffset, cImaginary + imaginaryOffset);
    }

    return {
        calculate: calculate
    }
}(specFactory);

var InputViewModel = new function(mandelbrotService, mandelbrotSpecCalculationService){

    var getZoomInPercentage = function(){ return 90; }
    var getDefaultSpecs = function()
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
    var getMandelbrotView = function(){ return document.getElementById("mandelbrotView");}

    var currentSpecs = getDefaultSpecs();

    var ClickOnMandelbrot = function (event)
    {
        console.log(currentSpecs);

        var e = event.target;
        var dim = e.getBoundingClientRect();
        var x = event.clientX - dim.left;
        var y = event.clientY - dim.top;

        mandelbrotService.getMandelbrot(currentSpecs);

        zoomedSpecs = mandelbrotSpecCalculationService.calculate(currentSpecs, x, y, getZoomInPercentage());

        currentSpecs = zoomedSpecs;
    };

    this.Init = function(){

        var defaultSpecs = getDefaultSpecs();
        mandelbrotService.openWebSocket(defaultSpecs, function(data){
            getMandelbrotView().src = "data:image/jpg;base64," + data;
        });

        getMandelbrotView().onclick = ClickOnMandelbrot;
    };

}(mandelbrotService, mandelbrotSpecCalculationService);