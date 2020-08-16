import _, { mat4, vec3, quat } from 'gl-matrix'

class Transformation {
    constructor( tx, ty, tz, scaleFactor, ryInDegrees ) {
        this.init(
            vec3.fromValues(tx, ty, tz),
            vec3.fromValues(scaleFactor, scaleFactor, scaleFactor),
            quat.fromValues(0.0, ryInDegrees * (Math.PI / 180.0), 0.0, 1.0)
        )
    }

    init( translation, scaling, rotation ) {
        this.translation = translation
        this.scaling = scaling
        this.rotation = rotation
    }

    setScaleFactor( scaleFactor ) {
        this.scaling = vec3.fromValues(scaleFactor, scaleFactor, scaleFactor)
    }

    setYRotation( degrees ) {
        this.rotation = quat.fromValues(0.0, degrees * (Math.PI / 180.0), 0.0, 1.0)
    }

    setTranslation( tx, ty, tz ) {
        this.translation = vec3.fromValues(tx, ty, tz)
    }

    translate( dx, dy, dz ) {
        vec3.add(this.translate, this.translate, vec3.fromValues(dx, dy, dz))
    }

    calculateTransformationMatrix(out) {
        mat4.fromRotationTranslationScale(out, this.rotation, this.translation, this.scaling)
    }
}

export { Transformation }