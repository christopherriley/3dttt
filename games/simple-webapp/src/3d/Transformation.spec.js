import { Transformation } from "./Transformation.js"

describe('default transformation', () => {
    var transformation = new Transformation(0.0, 0.0, 0.0, 1.0, 0.0);
    var result = [];
    transformation.calculateTransformationMatrix(result);

    it("returns expected default matrix", () => {
        expect(result).toEqual([
            1.0, 0.0, -0.0, 0.0,
            0.0, 1.0, 0.0, 0.0,
            0.0, 0.0, 1.0, 0.0,
            0.0, 0.0, 0.0, 1.0
        ]);
    });
});

describe('non-default transformation', () => {
    var transformation = new Transformation(2.0, 3.0, 4.0, 5.0, 6.0);
    var result = [];
    transformation.calculateTransformationMatrix(result);

    it("returns expected default matrix", () => {
        expect(result).toEqual([
            4.972609476841367, 0.0, -0.5226423163382674, 0.0,
            0.0, 5.0, 0.0, 0.0,
            0.5226423163382674, 0.0, 4.972609476841367, 0.0,
            12.035788219035803, 15.0, 18.845153274688933, 1.0
        ]);
    });
});
