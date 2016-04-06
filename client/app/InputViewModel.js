var InputViewModel = function(mandelbrotService, mandelbrotSpecCalculationService){
    if (!(this instanceof InputViewModel)) return new InputViewModel(andelbrotService, mandelbrotSpecCalculationService);

    var _getZoomInPercentage = function(){ return 90; }

    var _getDefaultSpecs = function()
    {
        return{
            width: _getMandelbrotView().offsetWidth,
            height: _getMandelbrotView().offsetHeight,
            iterations: 100,
            minR: -3,
            minI: -1.5,
            maxR: 1,
            maxI: 1.5
        }
    }
    var _currentSpecs;
    var _getMandelbrotView = function(){ return document.getElementById("mandelbrotView");}

    var _clickOnMandelbrot = function (event)
    {
        var e = event.target;
        var dim = e.getBoundingClientRect();
        var x = event.clientX - dim.left;
        var y = event.clientY - dim.top;


        _currentSpecs = mandelbrotSpecCalculationService.calculate(_currentSpecs, x, y, _getZoomInPercentage());
        mandelbrotService.getMandelbrot(_currentSpecs);

    };

    var init = function(){

        _currentSpecs = _getDefaultSpecs();
        mandelbrotService.openWebSocket(_currentSpecs, function(data){
            _getMandelbrotView().src = "data:image/jpg;base64," + data;
        });

        _getMandelbrotView().onclick = _clickOnMandelbrot;
    };

    return{
        init: init
    };
};