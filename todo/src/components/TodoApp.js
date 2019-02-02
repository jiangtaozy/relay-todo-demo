/*
 * Maintained by jemo from 2019.1.30 to now
 * Created by jemo on 2019.1.30 22:27
 * TodoApp
 */

import React, { Component } from 'react'
import {
  createFragmentContainer,
  graphql,
} from 'react-relay'
import TodoList from './TodoList'

class TodoApp extends Component {
  render() {
    return (
      <div>
        <TodoList viewer={this.props.viewer} />
      </div>
    )
  }
}

export default createFragmentContainer(TodoApp, {
  viewer: graphql`
    fragment TodoApp_viewer on User {
      id
      ...TodoList_viewer
    }
  `,
})
