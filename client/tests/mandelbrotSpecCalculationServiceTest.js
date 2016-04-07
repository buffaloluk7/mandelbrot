QUnit.test("iterationsAreCalculatedCorrect", function (assert) {
    var sut = mandelbrotSpecCalculationService;

    var mb0Specs = {
        width: 1000,
        height: 750,
        iterations: 100,
        minR: new BigNumber(-3),
        minI: new BigNumber(-1.5),
        maxR: new BigNumber(1),
        maxI: new BigNumber(1.5)
    };

    var firstRun = sut.calculate(mb0Specs, 500, 375, 90);
    var secondRun = sut.calculate(firstRun, 500, 375, 90);

    assert.equal(firstRun.iterations, 110);
    assert.equal(secondRun.iterations, 121);
});

QUnit.test("realAndImaginaryValuesAreCorrect", function (assert) {
    var sut = mandelbrotSpecCalculationService;

    var mb0Specs = {
        width: 1000,
        height: 750,
        iterations: 100,
        minR: new BigNumber(-3),
        minI: new BigNumber(-1.5),
        maxR: new BigNumber(1),
        maxI: new BigNumber(1.5)
    };

    // click is at -1 | 0
    var firstRun = sut.calculate(mb0Specs, 750, 375, 80);

    assert.floatEqual(firstRun.maxR, 0.8, 0.0001);
    assert.floatEqual(firstRun.minR, -2.4, 0.0001);
    assert.floatEqual(firstRun.maxI, 1.2, 0.0001);
    assert.floatEqual(firstRun.minI, 1.2, 0.0001);
});