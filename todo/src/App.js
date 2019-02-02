/*
 * Maintained by jemo from 2019.1.15 to now
 * Created by jemo on 2019.1.15 9:13
 * todo server
 */

import React, { Component } from 'react'
import { QueryRenderer, graphql } from 'react-relay'
import {
  Environment,
  Network,
  RecordSource,
  Store,
} from 'relay-runtime'
import TodoApp from './components/TodoApp'

function fetchQuery(operation, variables) {
  return fetch('http://192.168.31.11:3001/graphql', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      query: operation.text,
      variables,
    }),
  }).then(response => {
    return response.json();
  })
}

const modernEnvironment = new Environment({
  network: Network.create(fetchQuery),
  store: new Store(new RecordSource()),
})

class App extends Component {
  render() {
    return (
      <QueryRenderer
        environment={modernEnvironment}
        query={graphql`
          query AppQuery {
            viewer {
              ...TodoApp_viewer
            }
          }
        `}
        variables={{}}
        render={({error, props}) => {
          if(props) {
            return <TodoApp viewer={props.viewer} />
          } else {
            return <div>Loading</div>
          }
        }}
      />
    )
  }
}

export default App
