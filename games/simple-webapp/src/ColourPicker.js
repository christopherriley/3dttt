import React, { Component } from "react"

const Colour = {
    Red: 1,
    Blue: 2
}

function ColourSelectSquare(props) {
    return (
        <button className="colour-select-square" onClick={props.onClick}>
            {props.value}
        </button>
    )
}

class ColourPicker extends Component {
    render() {
        return (
            <div className="App">
                <h1>Please select your colour</h1>
                <ColourSelectSquare
                    value="Red"
                    onClick={() => this.props.cb(Colour.Red)}
                />
                <ColourSelectSquare
                    value="Blue"
                    onClick={() => this.props.cb(Colour.Blue)}
                />
            </div>
        )
    }
}

export { ColourPicker, Colour }
