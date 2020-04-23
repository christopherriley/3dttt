import React, { Component } from "react"
import { hot } from "react-hot-loader"
import "./App.css"
import { ColourPicker, Colour } from "./ColourPicker.js"
import { MoveFirstPicker } from "./MoveFirstPicker.js"
import { Game, NextAction, ActionResultStatus } from "./Game.js"

const NextMove = {
    RED_TO_MOVE: "RedToMove",
    BLUE_TO_MOVE: "BlueToMove"
}

class App extends Component {
    constructor(props) {
        super(props)
        this.state = {
            gameId: null,
            playerColour: null,
            nextAction: NextAction.START_NEW_GAME,
            moveFirst: null,
            boardState: null,
            playerLastPegClick: null,
            redScore: 0,
            blueScore: 0,
            error: null,
        }
    }


    renderMoveFirstPicker() {
        return (
            <MoveFirstPicker cb={moveFirst => this.handleMoveFirstClick(moveFirst)} />
        )
    }

    renderGame() {
        return (
            <Game
                gameId={this.state.gameId}
                action={this.state.nextAction}
                colour={this.state.playerColour}
                redScore={this.state.redScore}
                blueScore={this.state.blueScore}
                move_first={this.state.moveFirst}
                board_state={this.state.boardState}
                error={this.state.error}
                player_last_click={this.state.playerLastPegClick}
                actionResultCb={actionResult => this.handleActionResult(actionResult)}
                pegClickCb={peg => this.handlePegClick(peg)}
            />
        )
    }

    renderColourPicker() {
        return (
            <ColourPicker cb={colour => this.handleColourClick(colour)} />
        )
    }

    render() {
        if (this.state.playerColour == null) {
            return this.renderColourPicker()
        } else if (this.state.moveFirst == null) {
            return this.renderMoveFirstPicker()
        } else {
            return this.renderGame()
        }
    }

    handleColourClick(colour) {
        console.log("colour selected: " + colour)
        this.state.playerColour = colour
        this.setState(this.state)
    }

    handleMoveFirstClick(moveFirst) {
        console.log("move first: " + moveFirst)
        this.state.moveFirst = (moveFirst == "yes" ? true : false)
        this.state.nextMoveColour = (this.state.moveFirst ? this.state.playerColour : (this.state.playerColour == Colour.Red ? Colour.Blue : Colour.Red))
        this.setState(this.state)
    }

    handlePegClick(peg) {
        console.log("peg clicked: ", peg)
        if (this.state.nextAction == NextAction.PLAYER_TO_MOVE) {
            this.state.playerLastPegClick = peg
            this.state.nextAction = NextAction.PLAYER_MOVING
            this.setState(this.state)
        }
    }

    handleActionResult(actionResult) {
        if (actionResult.status == ActionResultStatus.SUCCESS) {
            if (this.state.nextAction == NextAction.START_NEW_GAME) {
                this.state.gameId = actionResult.id
                console.log("handleActionResult(): action was: NextAction.START_NEW_GAME, actionResult.NextMove: ", actionResult.nextMove)
                if (actionResult.nextMove == NextMove.RED_TO_MOVE && this.state.playerColour == Colour.Red ||
                    actionResult.nextMove == NextMove.BLUE_TO_MOVE && this.state.playerColour == Colour.Blue) {
                    this.state.nextAction = NextAction.PLAYER_TO_MOVE
                    console.log("changing state to: NextAction.PLAYER_TO_MOVE")
                }
                else if (actionResult.nextMove == NextMove.RED_TO_MOVE && this.state.playerColour == Colour.Blue ||
                    actionResult.nextMove == NextMove.BLUE_TO_MOVE && this.state.playerColour == Colour.Red) {
                    this.state.nextAction = NextAction.CPU_TO_MOVE
                    console.log("changing state to: NextAction.CPU_TO_MOVE")
                }

                this.state.boardState = actionResult.boardState
                this.state.redScore = actionResult.redScore
                this.state.blueScore = actionResult.blueScore

                this.setState(this.state)
            }
            else if (this.state.nextAction == NextAction.PLAYER_MOVING) {
                console.log("got success for PLAYER_MOVING state")
                this.state.boardState = actionResult.boardState
                this.state.nextAction = NextAction.CPU_TO_MOVE
                this.state.redScore = actionResult.redScore
                this.state.blueScore = actionResult.blueScore

                this.setState(this.state)
            }
            else if (this.state.nextAction == NextAction.CPU_TO_MOVE) {
                console.log("got success for CPU_TO_MOVE state")
                this.state.boardState = actionResult.boardState
                this.state.nextAction = NextAction.PLAYER_TO_MOVE
                this.state.redScore = actionResult.redScore
                this.state.blueScore = actionResult.blueScore

                this.setState(this.state)
            }
        } else {
            if (this.state.nextAction == NextAction.START_NEW_GAME) {
                this.state.nextAction = NextAction.START_NEW_GAME_FAILED
                this.state.error = actionResult.errorMsg

                this.setState(this.state)
            }
            else if (this.state.nextAction == NextAction.PLAYER_MOVING) {
                this.state.nextAction = NextAction.PLAYER_MOVING_FAILED
                this.state.error = actionResult.errorMsg

                this.setState(this.state)
            }
            else if (this.state.nextAction == NextAction.CPU_TO_MOVE) {
                this.state.nextAction = NextAction.CPU_TO_MOVE_FAILED
                this.state.error = actionResult.errorMsg

                this.setState(this.state)
            }
        }
    }
}

export default hot(module)(App)
