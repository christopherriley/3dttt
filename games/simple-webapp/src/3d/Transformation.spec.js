import { Transformation } from "./Transformation.js"

describe('default transformation', () => {
    var transformation = new Transformation(0.0, 0.0, 0.0, 1.0, 0.0);
    var result = [];
    transformation.calculateTransformationMatrix(result);

    it("returns expected default matrix", () => {
        expect(result).toEqual([
            1.0, 0.0, 0.0, 0.0,
            0.0, 1.0, 0.0, 0.0,
            0.0, 0.0, 1.0, 0.0,
            0.0, 0.0, 0.0, 1.0
        ]);
    });
});
