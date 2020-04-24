import React, { Component } from "react"

import { Board } from "./Board.js"
import { Colour } from "./ColourPicker.js"
import { Scoreboard } from "./Scoreboard.js"
import { GameError } from "./GameError.js"

import { PostCommand } from "./PostCommand"

const NextAction = {
    START_NEW_GAME: 1,
    START_NEW_GAME_FAILED: 2,
    PLAYER_TO_MOVE: 3,
    PLAYER_MOVING: 4,
    PLAYER_MOVING_FAILED: 5,
    CPU_TO_MOVE: 6,
    CPU_TO_MOVE_FAILED: 7
}

class Game extends Component {
    constructor(props) {
        super(props)
        this.url = this.props.url
        this.state = {
            gameId: null,
            playerColour: props.playerColour,
            nextAction: NextAction.START_NEW_GAME,
            moveFirst: props.moveFirst,
            nextMoveColour: (props.moveFirst ? props.playerColour : (props.playerColour == Colour.Red ? Colour.Blue : Colour.Red)),
            boardState: null,
            playerLastPegClick: null,
            redScore: 0,
            blueScore: 0,
            error: null,
        }
    }

    render() {
        if (this.state.nextAction == NextAction.START_NEW_GAME) {
            var params = {
                "colour": (this.state.playerColour == Colour.Red ? "red" : "blue"),
                "move_first": (this.state.moveFirst ? "TRUE" : "FALSE"),
            }

            var successNewState = (this.state.moveFirst ? NextAction.PLAYER_TO_MOVE : NextAction.CPU_TO_MOVE)

            new PostCommand(
                this.url,
                "newgame_1p",
                params,
                newGameState => this.handleCommandSuccess(newGameState, successNewState),
                error => this.handleCommandFail(error, NextAction.START_NEW_GAME_FAILED)).send()

            return (
                <h1>Start New Game</h1>
            )
        }
        else if (this.state.nextAction == NextAction.START_NEW_GAME_FAILED) {
            return (
                <GameError msg={this.state.error}/>
            )
        }
        else if (this.state.nextAction == NextAction.PLAYER_TO_MOVE) {
            return (
                <div>
                    <h1>Player to move</h1>
                    <Board
                        value={this.state.boardState}
                        cb={peg => this.handlePegClick(peg)}
                    />
                    <br/>
                    <br/>
                    <Scoreboard
                        redScore={this.state.redScore}
                        blueScore={this.state.blueScore}
                        playerColour={this.state.playerColour}
                    />
                </div>
            )
        }
        else if (this.state.nextAction == NextAction.PLAYER_MOVING) {
            var params = {
                "id": this.state.gameId,
                "peg": (this.state.playerLastPegClick),
            }

            new PostCommand(
                this.url,
                "move",
                params,
                newGameState => this.handleCommandSuccess(newGameState, NextAction.CPU_TO_MOVE),
                error => this.handleCommandFail(error, NextAction.PLAYER_MOVING_FAILED)).send()

            return (
                <div>
                    <h1>Player moving...</h1>
                    <Board
                        value={this.state.boardState}
                        cb={this.handlePegClick}
                    />
                    <br/>
                    <br/>
                    <Scoreboard
                        redScore={this.state.redScore}
                        blueScore={this.state.blueScore}
                        playerColour={this.state.playerColour}
                    />
                </div>
            )
        }
        else if (this.state.nextAction == NextAction.PLAYER_MOVING_FAILED) {
            return (
                <GameError msg={this.state.error}/>
            )
        }
        else if (this.state.nextAction == NextAction.CPU_TO_MOVE) {
            var params = {
                "id": this.state.gameId,
            }

            new PostCommand(
                this.url,
                "cpu_move",
                params,
                newGameState => this.handleCommandSuccess(newGameState, NextAction.PLAYER_TO_MOVE),
                error => this.handleCommandFail(error, NextAction.CPU_TO_MOVE_FAILED)).send()

            return (
                <div>
                    <h1>CPU to move</h1>
                    <Board
                        value={this.state.boardState}
                        cb={this.handlePegClick}
                    />
                    <br/>
                    <br/>
                    <Scoreboard
                        redScore={this.state.redScore}
                        blueScore={this.state.blueScore}
                        playerColour={this.state.playerColour}
                    />
                </div>
            )
        }
        else if (this.state.nextAction == NextAction.CPU_TO_MOVE_FAILED) {
            return (
                <GameError msg={this.state.error}/>
            )
        }
        else {
            return (
                <h1>Unknown Next Action</h1>
            )
        }
    }

    handlePegClick(peg) {
        console.log("peg clicked: ", peg)
        if (this.state.nextAction == NextAction.PLAYER_TO_MOVE) {
            this.state.playerLastPegClick = peg
            this.state.nextAction = NextAction.PLAYER_MOVING

            this.setState(this.state)
        }
    }

    handleCommandSuccess(newGameState, nextAction) {
        console.log("onCommandSuccesS(): nextAction: ", nextAction)
        this.state.error = null

        this.state.nextAction = nextAction

        if (!(newGameState.id === undefined)) {
            this.state.gameId = newGameState.id
            console.log("onCommandSuccess(): got a new game id: ", this.state.gameId)
        }
        this.state.boardState = newGameState.boardState
        this.state.redScore = newGameState.redScore
        this.state.blueScore = newGameState.blueScore

        this.setState(this.state)
    }

    handleCommandFail(error, nextAction) {
        this.state.error = error
        this.state.nextAction = nextAction

        this.setState(this.state)
    }
}

export { Game }
