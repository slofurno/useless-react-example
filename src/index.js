import React, { Component, PropTypes } from 'react'
import ReactDom from 'react-dom'

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {filter:"", list:[]}
    this.changeFilter = this.changeFilter.bind(this)
  }

  componentDidMount () {
    let self = this
    subs.push((n) => {
      let {list} = self.state
      list.push(n)
      self.setState({list})
    })
  }

  changeFilter (e) {
    this.setState({filter: e.target.value})
  }

  render () {

    let {list, filter} = this.state

    let filteredLog = list.filter(x => x.indexOf(filter) >= 0).map((x,i) => <li key={i}>{x}</li>)

    return(
        <div>
        <input type="text" value={this.state.filter} onChange={this.changeFilter}/>
       <ul>{filteredLog}</ul> 
        </div>
    )
  }
}

ReactDom.render(
  <App/>, document.getElementById("root")
)
