var InputViewModel = function (mandelbrotService, mandelbrotSpecCalculationService) {
    if (!(this instanceof InputViewModel)) return new InputViewModel(mandelbrotService, mandelbrotSpecCalculationService);

    var _getDefaultSpecs = function () {
        return {
            width: _mandelbrotView.offsetWidth,
            height: _mandelbrotView.offsetHeight,
            iterations: 100,
            minR: new BigNumber(-3),
            minI: new BigNumber(-1.5),
            maxR: new BigNumber(1),
            maxI: new BigNumber(1.5)
        }
    };

    var _currentSpecs;
    var _mandelbrotView = document.getElementById("mandelbrotView");

    var _wheelOnMandelbrot = function (event) {
        var e = event.target;
        var dim = e.getBoundingClientRect();
        var x = event.clientX - dim.left;
        var y = event.clientY - dim.top;

        var zoomIn = event.wheelDelta > 0;
        _currentSpecs = mandelbrotSpecCalculationService.calculate(_currentSpecs, x, y, zoomIn ? 90 : 110);

        mandelbrotService.sendMandelbrotCalculationRequest(_currentSpecs);
    };

    var init = function () {

        _currentSpecs = _getDefaultSpecs();
        mandelbrotService.openWebSocket(_currentSpecs, function (data) {
            _mandelbrotView.src = "data:image/jpg;base64," + data;
        });

        _mandelbrotView.onwheel = _wheelOnMandelbrot;
    };

    return {
        init: init
    };
};