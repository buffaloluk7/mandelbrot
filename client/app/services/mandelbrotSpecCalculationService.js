var mandelbrotSpecCalculationService = function(){

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
}();