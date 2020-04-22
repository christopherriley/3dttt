import React, { Component} from "react"

const styles = {
    slot: {
        textAlign: 'center'
    },
    slotHidden: {
        visibility: 'hidden'
    }
}

class Slot extends Component {
    render() {
        if (this.props.value === undefined) {
            return (
                <h3 className="slotHidden" style={styles.slotHidden}>-</h3>
            )
        }
        else {
            return (
                <h3 className="slot" style={styles.slot}>{this.props.value}</h3>
            )
        }
    }
}

export {Slot}
