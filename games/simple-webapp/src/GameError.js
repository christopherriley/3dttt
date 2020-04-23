import React, { Component } from "react"

const styles = {
    redSlot: {
        textAlign: 'center',
        color: 'red'
    },
    blueSlot: {
        textAlign: 'center',
        color: 'blue'
    },
    slotHidden: {
        visibility: 'hidden'
    }
}

class GameError extends Component {
    render() {
        var msg = "An unknown error has occurred"
        if (!(this.props.msg === undefined)) {
            msg = this.props.msg
        }
        return (
            <div>
                <h1>Uh oh</h1>
                <div/>
                <h2>{msg}</h2>
            </div>
        )
    }
}

export { GameError }
