import React, { Component} from "react"

import {Peg} from "./Peg.js"

class Board extends Component {
    render() {
        return (
            <div>
                <div>
                    <Peg value={this.props.value.Peg[0]}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[1]}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[2]}/>
                </div>
                <div>
                    <Peg/>
                    <Peg value={this.props.value.Peg[3]}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[4]}/>
                    <Peg/>
                </div>
                <div>
                    <Peg value={this.props.value.Peg[5]}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[6]}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[7]}/>
                </div>
            </div>
        )
    }
}

export {Board}
