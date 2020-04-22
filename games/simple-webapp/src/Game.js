import React, { Component} from "react"

import {Board} from "./Board.js"


const NextAction = {
    START_NEW_GAME: 1,
    PLAYER_TO_MOVE: 2
}

const ActionResultStatus = {
    SUCCESS: "success",
    FAIL: "fail"
}

class Game extends Component {
    postCommand(command, params) {
        var xhttp = new XMLHttpRequest()
        var url = 'http://localhost:8080/api/v1/game'

        xhttp.open('POST', url, true)
        xhttp.setRequestHeader('Content-type', 'application/json')

        var that = this
        xhttp.onreadystatechange = function() {
            if (this.readyState == 4) {
                if (this.status == 200) {
                    var jsonResponse = JSON.parse(xhttp.responseText)
                    that.props.actionResultCb({
                        status: ActionResultStatus.SUCCESS,
                        nextMove: jsonResponse.state.next_move,
                        boardState: jsonResponse.state.board_state
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
        if (this.props.action == NextAction.START_NEW_GAME) {
            var params = {
                "colour": this.props.colour,
                "move_first": (this.props.move_first ? "TRUE" : "FALSE"),
            }
            this.postCommand("newgame_1p", params)
            return(
                <h1>Start New Game</h1>
            )
        }
        else if (this.props.action == NextAction.PLAYER_TO_MOVE) {
            return (
                <div>
                    <h1>Player to move</h1>
                    <Board
                        value={this.props.board_state}
                        cb={this.props.pegClickCb}
                    />
                </div>
            )
        }
        else if (this.props.action == NextAction.CPU_TO_MOVE) {
            return(
                <div>
                    <h1>CPU to move</h1>
                    <Board
                        value={this.props.board_state}
                        cb={this.props.pegClickCb}
                    />
                </div>
            )
        }
        else {
            return(
                <h1>Unknown Next Action</h1>
            )
        }
    }
}

export {Game, NextAction, ActionResultStatus}
