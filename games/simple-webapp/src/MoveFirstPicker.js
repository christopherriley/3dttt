import React, { Component } from "react"


function MoveSelectSquare(props) {
    return (
        <button className="move-select-square" onClick={props.onClick}>
            {props.value}
        </button>
    )
}

class MoveFirstPicker extends Component {
    render() {
        return (
            <div className="App">
                <h1>Would you like to go first?</h1>
                <MoveSelectSquare
                    value="Yes"
                    onClick={() => this.props.cb("yes")}
                />
                <MoveSelectSquare
                    value="No"
                    onClick={() => this.props.cb("no")}
                />
            </div>
        )
    }
}

export { MoveFirstPicker }
