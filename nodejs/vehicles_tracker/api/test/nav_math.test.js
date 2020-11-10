const assert = require('assert');
const navMath = require('../src/nav_math');

describe('Distance testing', function() {
    it('zero distance', function() {
        assert.equal(0, navMath.distance(52.53, 13.403, 52.53, 13.403))
    });

    it('Kansas City and St. Louis distance', function() {
        assert.equal(38290005.03756017, navMath.distance(39.099912, -94.581213, 38.627089, -90.200203))
    });
});

describe('Navigation bearing testing', function() {
    it('Kansas City and St. Louis navigation bearing', function() {
        assert.equal(96.51262423499946, navMath.bearing(39.099912, -94.581213, 38.627089, -90.200203))
    });
});
