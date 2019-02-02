/*
 * Maintained by jemo from 2019.1.30 to now
 * Created by jemo on 2019.1.30 22:48
 * Todo
 */

import React, { Component } from 'react'
import {
  createFragmentContainer,
  graphql,
} from 'react-relay'
import ChangeTodoStatusMutation from '../mutations/ChangeTodoStatusMutation'

class Todo extends Component {
  handleCompleteChange = e => {
    const complete = e.target.checked
    ChangeTodoStatusMutation.commit(
      this.props.relay.environment,
      complete,
      this.props.todo,
      this.props.viewer,
    )
  }

  render() {
    return (
      <li>
        <div>
          <input
            checked={this.props.todo.complete}
            onChange={this.handleCompleteChange}
            type="checkbox"
          />
          <label>
            {this.props.todo.text}
          </label>
        </div>
      </li>
    )
  }
}

export default createFragmentContainer(Todo, {
  todo: graphql`
    fragment Todo_todo on Todo {
      complete
      id
      text
    }
  `,
  viewer: graphql`
    fragment Todo_viewer on User {
      id
    }
  `,
})
