import React, { Component } from "react"
import { Colour } from "./ColourPicker"

const styles = {
    container: {
        width: '120px'
    },
    redScore: {
        textAlign: 'right',
        color: 'red'
    },
    blueScore: {
        textAlign: 'right',
        color: 'blue'
    },
}

class Scoreboard extends Component {
    render() {
        if (this.props.playerColour == Colour.Blue) {
            return(
                <div className="container" style={styles.container}>
                    <div className="blueScore" style={styles.blueScore}>Human: {this.props.blueScore}</div>
                    <div className="redScore" style={styles.redScore}>CPU: {this.props.redScore}</div>
                </div>
            )
        }
        else {
            return(
                <div className="container" style={styles.container}>
                    <div className="redScore" style={styles.redScore}>Human: {this.props.redScore}</div>
                    <div className="blueScore" style={styles.blueScore}>CPU: {this.props.blueScore}</div>
                </div>
            )            
        }
    }
}

export { Scoreboard }
