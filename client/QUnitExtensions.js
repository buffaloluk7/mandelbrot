QUnit.extend(QUnit.assert, {
    floatEqual: function(actual, expected, margin, message){
        this.pushResult( {
            result: (actual - expected) < margin,
            actual: actual,
            expected: expected,
            message: message
        } );
    }
});