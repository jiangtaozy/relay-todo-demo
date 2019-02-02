/*
 * Maintained by jemo from 2019.1.27 to now
 * Created by jemo on 2019.1.27 10:44
 * database
 */

package graphql

type User struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Todos []*Todo `json:"todos"`
}

type Todo struct {
  Id string `json:"id"`
  Complete bool `json:"complete"`
  Text string `json:"text"`
}

var viewer = &User{
  Id: "1",
  Name: "Anonymous",
}

var todos = []*Todo{
  &Todo{"10", false, "todo0"},
  &Todo{"11", false, "todo1"},
  &Todo{"12", false, "todo2"},
}

func GetViewer() *User {
  return viewer
}

func GetUser(id string) *User {
  if id == viewer.Id {
    return viewer
  }
  return nil
}

func GetTodos() []*Todo {
  return todos
}

func GetTodo(id string) *Todo {
  for _, todo := range todos {
    if todo.Id == id {
      return todo
    }
  }
  return nil
}

func TodosToInterfaceSlice(todos ...*Todo) []interface{} {
  var interfaceSlice []interface{} = make([]interface{}, len(todos))
  for i, d := range todos {
    interfaceSlice[i] = d
  }
  return interfaceSlice
}

func ChangeTodoStatus(todoId string, complete bool) {
  for _, todo := range todos {
    if todoId == todo.Id {
      todo.Complete = complete
      break
    }
  }
}
