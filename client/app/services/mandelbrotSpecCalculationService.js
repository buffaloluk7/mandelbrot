var mandelbrotSpecCalculationService = function () {

    var calculate = function (oldSpecs, x, y, percentage) {

        var width = oldSpecs.maxR - oldSpecs.minR;
        var clickX = parseFloat(x) * (width / parseFloat(oldSpecs.width)) + oldSpecs.minR;
        var height = oldSpecs.maxI - oldSpecs.minI;
        var clickY = parseFloat(y) * (height / parseFloat(oldSpecs.height)) + oldSpecs.minI;

        percentage = (percentage > 1) ? percentage / 100 : percentage;

        var newWidth = width * percentage;
        var deltaWidth = (width - newWidth);

        var rightDistance = oldSpecs.maxR - clickX;
        var rightRelative = rightDistance / width;
        var rightOffset = deltaWidth * rightRelative;
        var leftOffset = (deltaWidth - rightOffset) * -1;


        var newHeight = height * percentage;
        var deltaHeight = height - newHeight;

        var upDistance = oldSpecs.maxI - clickY;
        var upRelative = upDistance / height;
        var upOffset = deltaHeight * upRelative;
        var downOffset = (deltaHeight - upOffset) * -1;

        var additionalIterations = (1 - percentage) * oldSpecs.iterations;

        var newSpec = {
            width: oldSpecs.width,
            height: oldSpecs.height,
            iterations: oldSpecs.iterations + additionalIterations,
            minR: oldSpecs.minR - leftOffset,
            maxR: oldSpecs.maxR - rightOffset,
            minI: oldSpecs.minI - downOffset,
            maxI: oldSpecs.maxI - upOffset
        };

        console.log(newSpec);

        return newSpec;
    };

    return {
        calculate: calculate
    }
}();