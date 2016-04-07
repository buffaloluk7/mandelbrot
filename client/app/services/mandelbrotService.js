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
        }
        webSocket.onclose = function (e) {
            console.log("DISCONNECTED");
        }
    }

    var getMandelbrot = function (specs) {
        if (webSocket == null) return;
        webSocket.send(_generateRequestArgument(specs));
    }

    var _generateRequestArgument = function (specs) {
        var nl = function (value) {
            return value + ";"
        };
        return nl(specs.width) + nl(specs.height) + nl(specs.minR) + nl(specs.minI) + nl(specs.maxR) + nl(specs.maxI) + nl(specs.iterations);
    };

    return {
        getMandelbrot: getMandelbrot,
        openWebSocket: openWebSocket
    };
}();