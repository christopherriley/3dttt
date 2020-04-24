class PostCommand {
    constructor(url, command, params, successCb, failCb) {
        this.url = url
        this.command = command
        this.params = params
        this.successCb = successCb
        this.failCb = failCb
    }

    send() {
        var xhttp = new XMLHttpRequest()
        //var url = 'http://localhost:8080/api/v1/game'

        xhttp.open('POST', this.url, true)
        xhttp.setRequestHeader('Content-type', 'application/json')

        var that = this
        xhttp.onreadystatechange = function () {
            if (this.readyState == 4) {
                var jsonResponse = JSON.parse(xhttp.responseText)
                if (this.status == 200) {
                    that.successCb({
                        id: jsonResponse.id,
                        boardState: jsonResponse.state.board_state,
                        redScore: jsonResponse.state.red_score,
                        blueScore: jsonResponse.state.blue_score,
                    })
                }
                else {
                    that.failCb(jsonResponse.error)
                }
            }
        }

        var body = JSON.stringify({
            "command": that.command,
            "params": that.params,
        })
        xhttp.send(body)        
    }
}

export { PostCommand }