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
                    <Peg value={this.props.value.Peg[0]} label='A' cb={this.props.cb}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[1]} label='B' cb={this.props.cb}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[2]} label='C' cb={this.props.cb}/>
                </div>
                <div className="rowDiv" style={styles.rowDiv}>
                    <Peg/>
                    <Peg value={this.props.value.Peg[3]} label='D' cb={this.props.cb}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[4]} label='E' cb={this.props.cb}/>
                    <Peg/>
                </div>
                <div className="rowDiv" style={styles.rowDiv}>
                    <Peg value={this.props.value.Peg[5]} label='F' cb={this.props.cb}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[6]} label='G' cb={this.props.cb}/>
                    <Peg/>
                    <Peg value={this.props.value.Peg[7]} label='H' cb={this.props.cb}/>
                </div>
            </div>
        )
    }
}

export {Board}
