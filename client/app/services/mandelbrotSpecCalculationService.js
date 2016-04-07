var mandelbrotSpecCalculationService = function () {

    var calculate = function (oldSpecs, x, y, percentage) {

        var widthComplex = oldSpecs.maxR.minus(oldSpecs.minR);
        var clickX = new BigNumber(x).times(widthComplex.dividedBy(oldSpecs.width)).plus(oldSpecs.minR);
        var heightComplex = oldSpecs.maxI.minus(oldSpecs.minI);
        var clickY = new BigNumber(y).times(heightComplex.dividedBy(oldSpecs.height)).plus(oldSpecs.minI);

        percentage = (percentage > 1) ? percentage / 100 : percentage;

        var newWidth = widthComplex.times(percentage);
        var deltaWidth = widthComplex.minus(newWidth);

        var rightDistance = oldSpecs.maxR.minus(clickX);
        var rightRelative = rightDistance.dividedBy(widthComplex);
        var rightOffset = deltaWidth.times(rightRelative);
        var leftOffset = deltaWidth.minus(rightOffset).times(-1);


        var newHeight = heightComplex.times(percentage);
        var deltaHeight = heightComplex.minus(newHeight);

        var upDistance = oldSpecs.maxI.minus(clickY);
        var upRelative = upDistance.dividedBy(heightComplex);
        var upOffset = deltaHeight.times(upRelative);
        var downOffset = deltaHeight.minus(upOffset).times(-1);

        var additionalIterations = (1 - percentage) * oldSpecs.iterations;

        var newSpec = {
            width: oldSpecs.width,
            height: oldSpecs.height,
            iterations: oldSpecs.iterations + additionalIterations,
            minR: oldSpecs.minR.minus(leftOffset),
            maxR: oldSpecs.maxR.minus(rightOffset),
            minI: oldSpecs.minI.minus(downOffset),
            maxI: oldSpecs.maxI.minus(upOffset)
        };

        return newSpec;
    };

    return {
        calculate: calculate
    }
}();