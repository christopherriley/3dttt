import { Transformation } from "./Transformation.js"

describe('calculateTransformationMatrix test', () => {
    var result;
    var transformation;
    var tx, ty, tz;
    var scaleFactor;
    var yRotationInDegrees;
    var expectedMatrix;

    beforeEach( () => {
        transformation = new Transformation(tx, ty, tz, scaleFactor, yRotationInDegrees);
        result = [];
        transformation.calculateTransformationMatrix(result);
    });

    describe('default transformation', () => {
        tx = 0.0;
        ty = 0.0;
        tz = 0.0;
        scaleFactor = 1.0;
        yRotationInDegrees = 0.0;

        expectedMatrix = [
            1.0, 0.0, -0.0, 0.0,
            0.0, 1.0, 0.0, 0.0,
            0.0, 0.0, 1.0, 0.0,
            0.0, 0.0, 0.0, 1.0
        ];

        it("returns expected matrix", () => {
            expect(result).toEqual(expectedMatrix);
        });
    });

    describe('non-default transformation', () => {
        tx = 2.0;
        ty = 3.0;
        tz = 4.0;
        scaleFactor = 5.0;
        yRotationInDegrees = 6.0;

        expectedMatrix = [
            4.972609476841367, 0.0, -0.5226423163382674, 0.0,
            0.0, 5.0, 0.0, 0.0,
            0.5226423163382674, 0.0, 4.972609476841367, 0.0,
            12.035788219035803, 15.0, 18.845153274688933, 1.0
        ];

        it("returns expected matrix", () => {
            expect(result).toEqual(expectedMatrix);
        });
    });
});
