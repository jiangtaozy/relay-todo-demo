A relay demo

## Start

      go run todo.go
      cd todo
      yarn relay
      yarn start

## Step

      create-react-app init todo
      cd todo
      yarn run eject
      yarn add react-relay
      yarn add relay-compiler babel-plugin-relay graphql --dev
      Edit package.json
      "scripts": {
        "relay": "relay-compiler --src ./src --schema ./schema.graphql"
      },
      "babel": {
        "presets": [
          "react-app"
        ],
        "plugins": [
          "relay"
        ]
      },

## Prerequisites

- go version 1.11

## Dependencies

- react-relay
- babel-plugin-relay
- relay-compiler
- graphql
- create-react-app
- github.com/graphql-go/graphql
- github.com/graphql-go/relay
- github.com/rs/cors

## Links

- https://facebook.github.io/relay/docs/en/installation-and-setup.html
- https://www.prisma.io/blog/getting-started-with-relay-modern-46f8de6bd6ec/
- https://github.com/relayjs/relay-examples/tree/master/todo
- https://github.com/sogko/golang-relay-starter-kit
- https://github.com/graphql-go/relay/tree/master/examples/starwars

## License

MIT licensed
