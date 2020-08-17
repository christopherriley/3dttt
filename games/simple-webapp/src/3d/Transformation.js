import _, { mat4, vec3, quat } from 'gl-matrix'

class Transformation {
    constructor( tx, ty, tz, scaleFactor, ryInDegrees ) {
        this.translation = vec3.fromValues(tx, ty, tz);
        this.scaling = vec3.fromValues(scaleFactor, scaleFactor, scaleFactor);
        //this.rotation = quat.fromValues(0.0, ryInDegrees * (Math.PI / 180.0), 0.0, 1.0);
        this.yRotationInRads = ryInDegrees * (Math.PI / 180.0);
    }

    setScaleFactor( scaleFactor ) {
        this.scaling = vec3.fromValues(scaleFactor, scaleFactor, scaleFactor)
    }

    setYRotation( degrees ) {
        this.yRotationInRads = degrees * (Math.PI / 180.0);
    }

    setTranslation( tx, ty, tz ) {
        this.translation = vec3.fromValues(tx, ty, tz)
    }

    translate( dx, dy, dz ) {
        vec3.add(this.translate, this.translate, vec3.fromValues(dx, dy, dz))
    }

    calculateTransformationMatrix(out) {
        //mat4.fromRotationTranslationScale(out, this.rotation, this.translation, this.scaling)
        mat4.fromYRotation(out, this.yRotationInRads);
        mat4.scale(out, out, this.scaling);
        mat4.translate(out, out, this.translation);
    }
}

export { Transformation }