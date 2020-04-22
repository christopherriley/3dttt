import React, { Component} from "react"

import {Slot} from "./Slot.js"

class Peg extends Component {
    render() {
        if (this.props.value === undefined) {
            return (
                <table style={{display: "inline-block"}}>
                    <tbody>
                        <tr>
                            <td><Slot/></td>
                        </tr>
                        <tr>
                            <td><Slot/></td>
                        </tr>
                        <tr>
                            <td><Slot/></td>
                        </tr>
                    </tbody>
                </table>
            )
        }
        else {
            return (
                <table style={{display: "inline-block"}}>
                    <tbody>
                        <tr>
                            <td><Slot value={this.props.value.Slot[2]}/></td>
                        </tr>
                        <tr>
                            <td><Slot value={this.props.value.Slot[1]}/></td>
                        </tr>
                        <tr>
                            <td><Slot value={this.props.value.Slot[0]}/></td>
                        </tr>
                    </tbody>
                </table>
            )
        }
    }
}

export {Peg}
