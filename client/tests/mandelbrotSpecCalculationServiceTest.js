QUnit.test( "iterationsAreCalculatedCorrect", function( assert ) {
    var sut = new MandelbrotSpecCalculationService ();

    mb0Specs =  {
                    width: 1000,
                    height: 750,
                    iterations: 100,
                    minR: -3,
                    minI: -1.5,
                    maxR: 1,
                    maxI: 1.5
                }

    var firstRun = sut.Calculate(mb0Specs, 5000, 375, 90);
    var secondRun = sut.Calculate(firstRun, 5000, 375, 90);

    assert.ok(parseInt(firstRun.iterations) === 110);
    assert.ok(parseInt(secondRun.iterations) === 121);
});

QUnit.test( "realAndImaginaryValuesAreCorrect", function( assert ) {
    var sut = new MandelbrotSpecCalculationService ();

    mb0Specs =  {
                    width: 1000,
                    height: 750,
                    iterations: 100,
                    minR: -3,
                    minI: -1.5,
                    maxR: 1,
                    maxI: 1.5
                }

    var firstRun = sut.Calculate(mb0Specs, 5000, 375, 90);
    var secondRun = sut.Calculate(firstRun, 5000, 375, 90);

    assert.ok(parseInt(firstRun.iterations) === 110);
    assert.ok(parseInt(secondRun.iterations) === 121);
});