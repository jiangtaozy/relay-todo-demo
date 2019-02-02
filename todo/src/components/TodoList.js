/*
 * Maintained by jemo from 2019.1.30 to now
 * Created by jemo on 2019.1.30 22:33
 * TodoList
 */

import React, { Component } from 'react'
import {
  createFragmentContainer,
  graphql,
} from 'react-relay'
import Todo from './Todo'

class TodoList extends Component {
  renderTodos() {
    return this.props.viewer.todos.edges.map(edge => (
      <Todo key={edge.node.id}
        todo={edge.node}
        viewer={this.props.viewer}
      />
    ))
  }
  render() {
    return (
      <section>
        <ul>
          {this.renderTodos()}
        </ul>
      </section>
    )
  }
}

export default createFragmentContainer(TodoList, {
  viewer: graphql`
    fragment TodoList_viewer on User {
      todos(
        first: 2147483647 # max GraphQLInt
      ) @connection(key: "TodoList_todos") {
        edges {
          node {
            id
            complete
            ...Todo_todo
          }
        }
      }
      id
      ...Todo_viewer
    }
  `,
})
