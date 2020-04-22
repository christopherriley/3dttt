import React, { Component} from "react"


const NextAction = {
    START_NEW_GAME: 1
  }

class Game extends Component {
    postCommand(command, params) {
        var http = new XMLHttpRequest()
        var url = 'http://localhost:8080/api/v1/game'

        http.open('POST', url, true)
        http.setRequestHeader('Content-type', 'application/json')

        var body = JSON.stringify({
            "command": command,
            "params": params,
        })
        http.send(body)
    }
    render() {
        if (this.props.action == NextAction.START_NEW_GAME) {
            var params = {
                "colour": "red",
                "move_first": "TRUE",
            }
            this.postCommand("newgame_1p", params)
            return(
                <h1>Start New Game</h1>
            )
        }
        else {
            return(
                <h1>Unknown Next Action</h1>
            )
        }
    }
}

export {Game, NextAction}
