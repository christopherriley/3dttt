import React, { Component} from "react"

import {Peg} from "./Peg.js"

const styles = {
    containerDiv: {
        border: '0px solid black'
    },
    rowDiv: {
        border: '0px solid black'
    }
}

class Board extends Component {
    render() {
        return (
            <div className="containerDiv" style={styles.containerDiv}>
                <div className="rowDiv" style={styles.rowDiv}>
                    <Peg value={this.props.value.Peg[0]}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[1]}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[2]}/>
                </div>
                <div className="rowDiv" style={styles.rowDiv}>
                    <Peg/>
                    <Peg value={this.props.value.Peg[3]}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[4]}/>
                    <Peg/>
                </div>
                <div className="rowDiv" style={styles.rowDiv}>
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
