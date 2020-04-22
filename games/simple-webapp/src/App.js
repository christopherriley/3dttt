import React, { Component} from "react"
import {hot} from "react-hot-loader"
import "./App.css"
import {ColourPicker, Colour} from "./ColourPicker.js"
import {MoveFirstPicker} from "./MoveFirstPicker.js"
import {Game, NextAction} from "./Game.js"

class App extends Component{
  constructor(props) {
    super(props)
    this.state = {
      playerColour: null,
      nextAction: NextAction.START_NEW_GAME,
      moveFirst: null,
    }
  }


  renderMoveFirstPicker() {
    return(
      <MoveFirstPicker cb={moveFirst => this.handleMoveFirstClick(moveFirst)}/>
    )
  }

  renderGame(nextAction) {
    return (
      <Game action={nextAction} cb={move => this.handleMoveClick(move)}/>
    )
  }

  renderColourPicker() {
    return (
      <ColourPicker cb={colour => this.handleColourClick(colour)}/>
    )
  }

  render() {
    console.log("App.render()")
    if (this.state.playerColour == null) {
      return this.renderColourPicker()
    } else if (this.state.moveFirst == null) {
      return this.renderMoveFirstPicker()
    } else {
      return this.renderGame(this.state.nextAction)
    }
  }

  handleColourClick(colour) {
    console.log("colour selected: " + colour)
    this.state.playerColour = (colour == "red" ? Colour.Red : Colour.Blue)
    this.setState(this.state)
  }

  handleMoveFirstClick(moveFirst) {
    console.log("move first: " + moveFirst)
    this.state.moveFirst = (moveFirst == "yes" ? true : false)
    this.state.nextMoveColour = (this.state.moveFirst ? this.state.playerColour : (this.state.playerColour == Colour.Red ? Colour.Blue : Colour.Red))
    this.setState(this.state)
  }

  handleMoveClick(peg) {
    console.log("move: " + peg)
  }
}

export default hot(module)(App)
