import React, { Component} from "react"
import {hot} from "react-hot-loader"
import "./App.css"

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

function MoveSelectSquare(props) {
  return (
    <button className="move-select-square" onClick={props.onClick}>
      {props.value}
    </button>
  )
}

class App extends Component{
  constructor(props) {
    super(props)
    this.state = {
      playerColour: null,
      moveFirst: null,
    }
  }

  renderColourPicker() {
    return(
      <div className="App">
      <h1>Please select your colour</h1>
      <ColourSelectSquare
        value="Red"
        onClick={() => this.handleColourClick("red")}
      />
      <ColourSelectSquare
        value="Blue"
        onClick={() => this.handleColourClick("blue")}
      />
    </div>
    )
  }

  renderMoveFirstPicker() {
    return(
      <div className="App">
        <h1>Would you like to go first?</h1>
        <MoveSelectSquare
        value="Yes"
        onClick={() => this.handleMoveFirstClick("yes")}
      />
        <MoveSelectSquare
        value="No"
        onClick={() => this.handleMoveFirstClick("no")}
      />
      </div>
    )
  }

  renderGame() {
    return(
      <h1>suuuuup</h1>
    )
  }

  render(){
    console.log("App.render()")
    if (this.state.playerColour == null) {
      return this.renderColourPicker()
    } else if (this.state.moveFirst == null) {
      return this.renderMoveFirstPicker()
    } else {
      return this.renderGame()
    }
  }

  handleColourClick(colour){
    console.log("colour selected: " + colour)
    this.state.playerColour = (colour == "red" ? Colour.Red : Colour.Blue)
    this.setState(this.state)
  }

  handleMoveFirstClick(moveFirst){
    console.log("move first: " + moveFirst)
    this.state.moveFirst = (moveFirst == "yes" ? true : false)
    this.setState(this.state)
  }
}

export default hot(module)(App)
