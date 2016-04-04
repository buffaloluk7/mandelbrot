QUnit.test( "zoom into middle of the picture", function( assert ) {
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

    var specs = sut.Calculate(mb0Specs, 5000, 375, 10);

    console.log(specs);
});