import React, { Component } from "react"
import { Colour } from "./ColourPicker"

const styles = {
    redSlot: {
        textAlign: 'center',
        color: 'red'
    },
    blueSlot: {
        textAlign: 'center',
        color: 'blue'
    },
    slotHidden: {
        visibility: 'hidden'
    }
}

class Slot extends Component {
    render() {
        if (this.props.value == Colour.Red) {
            return (
                <h3 className="redSlot" style={styles.redSlot}>X</h3>
            )
        }
        else if (this.props.value == Colour.Blue) {
            return (
                <h3 className="blueSlot" style={styles.blueSlot}>X</h3>
            )
        }
        else {
            return (
                <h3 className="slotHidden" style={styles.slotHidden}>-</h3>
            )
        }
    }
}

export { Slot }
