function postCommand(url, command, params) {
    var xhttp = new XMLHttpRequest()

    return new Promise(function(resolve, reject) {
        xhttp.onreadystatechange = function () {
            if (xhttp.readyState == 4) {
                var jsonResponse = JSON.parse(xhttp.responseText)
                if (xhttp.status == 200) {
                    resolve({
                        id: jsonResponse.id,
                        boardState: jsonResponse.state.board_state,
                        redScore: jsonResponse.state.red_score,
                        blueScore: jsonResponse.state.blue_score,
                    })
                }
                else {
                    reject({
                        error: jsonResponse.error,
                    })
                }
            }
        }

        xhttp.open('POST', url, true)
        xhttp.setRequestHeader('Content-type', 'application/json')

        var body = JSON.stringify({
            "command": command,
            "params": params,
        })

        xhttp.send(body)
    })
}

export { postCommand }
