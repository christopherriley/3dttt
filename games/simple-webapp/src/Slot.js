import React, { Component} from "react"

const styles = {
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
                <h3>{this.props.value}</h3>
            )
        }
    }
}

export {Slot}
