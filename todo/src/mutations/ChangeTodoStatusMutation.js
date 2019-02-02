/*
 * Maintained by jemo from 2019.2.2 to now
 * Created by jemo on 2019.2.2 10:01
 * Change todo status mutation
 */

import {
  commitMutation,
  graphql,
} from 'react-relay'

const mutation = graphql`
  mutation ChangeTodoStatusMutation($input: ChangeTodoStatusInput!) {
    changeTodoStatus(input: $input) {
      todo {
        id
        complete
      }
      viewer {
        id
      }
    }
  }
`

function getOptimisticResponse(complete, todo, user) {
  const viewerPayload = {id: user.id}
  return {
    changeTodoStatus: {
      todo: {
        complete: complete,
        id: todo.id,
      },
      viewer: viewerPayload,
    }
  }
}

function commit (environment, complete, todo, user) {
  return commitMutation(environment, {
    mutation,
    variables: {
      input: {
        complete,
        id: todo.id,
        clientMutationId: todo.id,
      },
    },
    optimisticResponse: getOptimisticResponse(complete, todo, user),
  })
}

export default {commit}
