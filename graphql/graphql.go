/*
 * Maintained by jemo from 2019.1.15 to now
 * Created by jemo on 2019.1.15 10:14
 * graphql
 */

package graphql

import (
  "log"
  "time"
  "net/http"
  "encoding/json"
  "golang.org/x/net/context"
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/relay"
  "github.com/dancannon/gorethink"
)

const (
  dialTimeout = 2 * time.Second
  requestTimeout = 10 * time.Second
)

// rethink session
var session *gorethink.Session

var userType *graphql.Object
var todoType *graphql.Object
var nodeDefinitions *relay.NodeDefinitions
var todoConnection *relay.GraphQLConnectionDefinitions
var Schema graphql.Schema

func Init() {
  /*
  session, err := gorethink.Connect(gorethink.ConnectOpts{
    Address: "localhost:28015",
    Database: "todo",
    MaxOpen: 40,
  })
  if err != nil {
    log.Fatalln("initDbError, err: ", err)
  }
  */
  nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
    IDFetcher: func(id string, info graphql.ResolveInfo, ct context.Context) (interface{}, error) {
      resolvedID := relay.FromGlobalID(id)
      if resolvedID.Type == "User" {
        return GetUser(resolvedID.ID), nil
      }
      if resolvedID.Type == "Todo" {
        return GetTodo(resolvedID.ID), nil
      }
      return nil, nil
    },
    TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
      switch p.Value.(type) {
      case *User:
        return userType
      case *Todo:
        return todoType
      }
      return nil
    },
  })
  todoType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Todo",
    Description: "todo",
    Fields: graphql.Fields{
      "id": relay.GlobalIDField("Todo", nil),
      "complete": &graphql.Field{
        Description: "todo complete",
        Type: graphql.Boolean,
      },
      "text": &graphql.Field{
        Description: "todo text",
        Type: graphql.String,
      },
    },
    Interfaces: []*graphql.Interface{
      nodeDefinitions.NodeInterface,
    },
  })
  todoConnection = relay.ConnectionDefinitions(relay.ConnectionConfig{
    Name: "TodoConnection",
    NodeType: todoType,
  })
  userType = graphql.NewObject(graphql.ObjectConfig{
    Name: "User",
    Fields: graphql.Fields{
      "id": relay.GlobalIDField("user", nil),
      "todos": &graphql.Field{
        Type: todoConnection.ConnectionType,
        Description: "todo list",
        Args: relay.ConnectionArgs,
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          args := relay.NewConnectionArguments(p.Args)
          dataSlice := TodosToInterfaceSlice(GetTodos()...)
          return relay.ConnectionFromArray(dataSlice, args), nil
        },
      },
    },
    Interfaces: []*graphql.Interface{
      nodeDefinitions.NodeInterface,
    },
  })
  rootQuery := graphql.NewObject(graphql.ObjectConfig{
    Name: "RootQuery",
    Fields: graphql.Fields{
      "viewer": &graphql.Field{
        Type: userType,
        Description: "user info",
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return GetViewer(), nil
        },
      },
      "node": nodeDefinitions.NodeField,
    },
  })
  Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: rootQuery,
  })
}

type PostData struct {
  Query string `json:"query"`
  Variables map[string]interface{} `json:"variables"`
}

func GraphqlHandle(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  var data PostData
  err := decoder.Decode(&data)
  if err != nil {
    log.Println("GraphqlHandleDecodeError, err: ", err)
    panic(err)
  }
  res := graphql.Do(graphql.Params{
    Schema: Schema,
    RequestString: data.Query,
    VariableValues: data.Variables,
  })
  if len(res.Errors) > 0 {
    log.Printf("GraphqlHandleResError, res.Errors: %v\n", res.Errors)
  }
  json.NewEncoder(w).Encode(res)
}
