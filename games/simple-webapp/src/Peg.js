import React, { Component} from "react"

import {Slot} from "./Slot.js"

class Peg extends Component {
    render() {
        if (this.props.value === undefined) {
            return (
                <div style={{display: "inline-block"}}>
                    <div><Slot/></div>
                    <div><Slot/></div>
                    <div><Slot/></div>
                </div>
            )
        }
        else {
            return (
                <div style={{display: "inline-block"}}>
                    <div><Slot value={this.props.value.Slot[2]}/></div>
                    <div><Slot value={this.props.value.Slot[1]}/></div>
                    <div><Slot value={this.props.value.Slot[0]}/></div>
                </div>
            )
        }
    }
}

export {Peg}
