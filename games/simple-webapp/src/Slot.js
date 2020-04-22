import React, { Component} from "react"

class Slot extends Component {
    render() {
        if (this.props.value === undefined) {
            return (
                <h3 style={{visibility: "hidden"}}>-</h3>
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
