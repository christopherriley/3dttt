import React, { Component } from "react"
import { hot } from "react-hot-loader"
import "./App.css"
import { ColourPicker } from "./ColourPicker.js"
import { MoveFirstPicker } from "./MoveFirstPicker.js"
import { Game } from "./Game.js"

const NextAction = {
    PICK_COLOUR: 1,
    PICK_FIRST_MOVE: 2,
    PLAY_GAME: 3
}

class App extends Component {
    constructor(props) {
        super(props)
        this.state = {
            nextAction: NextAction.PICK_COLOUR,
            playerColour: null,
            moveFirst: null
        }
    }

    render() {
        if (this.state.nextAction == NextAction.PICK_COLOUR) {
            return (
                <ColourPicker cb={colour => this.handleColourClick(colour)} />
            )
        }
        else if (this.state.nextAction == NextAction.PICK_FIRST_MOVE) {
            return (
                <MoveFirstPicker cb={moveFirst => this.handleMoveFirstClick(moveFirst)} />
            )
        }
        else {
            return (
                <Game
                    url = 'http://localhost:8080/api/v1/game'
                    playerColour={this.state.playerColour}
                    moveFirst={this.state.moveFirst}
                />
            )
        }
    }

    handleColourClick(colour) {
        console.log("colour selected: " + colour)
        this.state.playerColour = colour
        this.state.nextAction = NextAction.PICK_FIRST_MOVE

        this.setState(this.state)
    }

    handleMoveFirstClick(moveFirst) {
        console.log("move first: " + moveFirst)
        this.state.moveFirst = (moveFirst == "yes" ? true : false)
        this.state.nextAction = NextAction.PLAY_GAME

        this.setState(this.state)
    }
}

export default hot(module)(App)
