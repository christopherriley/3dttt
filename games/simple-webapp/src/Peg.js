import React, { Component} from "react"

import {Slot} from "./Slot.js"

const styles = {
    pegDiv: {
        display: 'inline-block',
        border: '1px solid black',
        width: '20px'
    },
    spacerDiv: {
        display: 'inline-block',
        width: '20px'
    }
}

class Peg extends Component {
    render() {
        if (this.props.value === undefined) {
            return (
                <div className="spacerDiv" style={styles.spacerDiv}>
                    <div><Slot/></div>
                    <div><Slot/></div>
                    <div><Slot/></div>
                </div>
            )
        }
        else {
            return (
                <div className="pegDiv" style={styles.pegDiv}>
                    <div><Slot value={this.props.value.Slot[2]}/></div>
                    <div><Slot value={this.props.value.Slot[1]}/></div>
                    <div><Slot value={this.props.value.Slot[0]}/></div>
                </div>
            )
        }
    }
}

export {Peg}
