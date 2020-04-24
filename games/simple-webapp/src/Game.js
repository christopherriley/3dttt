import React, { Component } from "react"

import { Board } from "./Board.js"
import { Colour } from "./ColourPicker.js"
import { Scoreboard } from "./Scoreboard.js"
import { GameError } from "./GameError.js"


const NextAction = {
    START_NEW_GAME: 1,
    START_NEW_GAME_FAILED: 2,
    PLAYER_TO_MOVE: 3,
    PLAYER_MOVING: 4,
    PLAYER_MOVING_FAILED: 5,
    CPU_TO_MOVE: 6,
    CPU_TO_MOVE_FAILED: 7
}

const NextMove = {
    RED_TO_MOVE: "RedToMove",
    BLUE_TO_MOVE: "BlueToMove"
}

const ActionResultStatus = {
    SUCCESS: "success",
    FAIL: "fail"
}

class Game extends Component {
    constructor(props) {
        super(props)
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

    postCommand(command, params, cb) {
        var xhttp = new XMLHttpRequest()
        var url = 'http://localhost:8080/api/v1/game'

        xhttp.open('POST', url, true)
        xhttp.setRequestHeader('Content-type', 'application/json')

        xhttp.onreadystatechange = function () {
            if (this.readyState == 4) {
                if (this.status == 200) {
                    var jsonResponse = JSON.parse(xhttp.responseText)
                    cb({
                        status: ActionResultStatus.SUCCESS,
                        id: jsonResponse.id,
                        nextMove: jsonResponse.state.next_move,
                        boardState: jsonResponse.state.board_state,
                        redScore: jsonResponse.state.red_score,
                        blueScore: jsonResponse.state.blue_score,
                    })
                }
                else {
                    var jsonResponse = JSON.parse(xhttp.responseText)
                    cb({
                        status: ActionResultStatus.FAIL,
                        errorMsg: jsonResponse.error,
                    })
                }
            }
        }

        var body = JSON.stringify({
            "command": command,
            "params": params,
        })
        xhttp.send(body)
    }
    render() {
        if (this.state.nextAction == NextAction.START_NEW_GAME) {
            var params = {
                "colour": (this.state.playerColour == Colour.Red ? "red" : "blue"),
                "move_first": (this.state.moveFirst ? "TRUE" : "FALSE"),
            }
            this.postCommand("newgame_1p", params, result => this.handlePostCommandResult(result))
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
            this.postCommand("move", params, result => this.handlePostCommandResult(result))
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
            this.postCommand("cpu_move", params, result => this.handlePostCommandResult(result))
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

    handlePostCommandResult(postCommandResult) {
        if (postCommandResult.status == ActionResultStatus.SUCCESS) {
            this.state.boardState = postCommandResult.boardState
            this.state.redScore = postCommandResult.redScore
            this.state.blueScore = postCommandResult.blueScore
            this.state.error = null

            if (this.state.nextAction == NextAction.START_NEW_GAME) {
                this.state.gameId = postCommandResult.id

                if (postCommandResult.nextMove == NextMove.RED_TO_MOVE && this.state.playerColour == Colour.Red ||
                    postCommandResult.nextMove == NextMove.BLUE_TO_MOVE && this.state.playerColour == Colour.Blue) {
                    this.state.nextAction = NextAction.PLAYER_TO_MOVE
                }
                else {
                    this.state.nextAction = NextAction.CPU_TO_MOVE
                }
            }
            else if (this.state.nextAction == NextAction.PLAYER_MOVING) {
                this.state.nextAction = NextAction.CPU_TO_MOVE
            }
            else if (this.state.nextAction == NextAction.CPU_TO_MOVE) {
                this.state.nextAction = NextAction.PLAYER_TO_MOVE
            }
        } else {
            this.state.error = postCommandResult.errorMsg

            if (this.state.nextAction == NextAction.START_NEW_GAME) {
                this.state.nextAction = NextAction.START_NEW_GAME_FAILED
            }
            else if (this.state.nextAction == NextAction.PLAYER_MOVING) {
                this.state.nextAction = NextAction.PLAYER_MOVING_FAILED
            }
            else if (this.state.nextAction == NextAction.CPU_TO_MOVE) {
                this.state.nextAction = NextAction.CPU_TO_MOVE_FAILED
            }
        }

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
}

export { Game }
