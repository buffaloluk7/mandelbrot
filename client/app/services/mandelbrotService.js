var mandelbrotService = function () {

    var webSocket = null;

    var openWebSocket = function (initialSpecs, callback) {
        webSocket = new WebSocket('ws://localhost:8080/mandelbrot', []);

        webSocket.onerror = function (error) {
            console.log('WebSocket Error ' + error);
        };
        webSocket.onmessage = function (e) {
            callback(e.data);
        };
        webSocket.onopen = function (e) {
            var webSocketArguments = _generateRequestArgument(initialSpecs);
            console.log(webSocketArguments);
            webSocket.send(webSocketArguments);
        };
        webSocket.onclose = function (e) {
            console.log("DISCONNECTED");
        }
    };

    var sendMandelbrotCalculationRequest = function (specs) {
        if (webSocket == null) return;
        var specsAsString = _generateRequestArgument(specs);

        console.log(specsAsString);

        webSocket.send(specsAsString);
    };

    var _generateRequestArgument = function (specs) {
        var nl = function (value) {
            return value + ";"
        };
        return nl(specs.width) + nl(specs.height) + nl(specs.minR.toFixed(10)) + nl(specs.minI.toFixed(10)) + nl(specs.maxR.toFixed(10)) + nl(specs.maxI.toFixed(10)) + nl(parseInt(specs.iterations));
    };

    return {
        sendMandelbrotCalculationRequest: sendMandelbrotCalculationRequest,
        openWebSocket: openWebSocket
    };
}();