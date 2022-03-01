import React, { useEffect, useReducer, useState } from "react";
import "./App.css";
import { DragDropContext, DropResult } from "react-beautiful-dnd";
import InputField from "./Components/InputField";
import TodoList from "./Components/TodoList";
import { Todo, ActionsType, Actions } from "./model";
import axios from "./axios";

const TodoReducer = (todos: Todo[], action: Actions) => {
  switch (action.type) {
    case ActionsType.GET:
      if (action.payload) {
        return Object.entries(action.payload).map(([_, todo]) => {
          return {
            id: todo.todoid,
            todo: todo.todo,
            deadline: new Date(todo.deadline),
            isDone: todo.isdone,
          };
        });
      } else return todos;
    case ActionsType.ADD:
      const data = {
        id: Date.now(),
        todo: action.payload.todo,
        deadline: action.payload.deadline,
        isDone: false,
      };
      console.log(JSON.stringify(data));
      axios.post("api/todo", data).then((response) => console.log(response));
      return [...todos, data];
    case ActionsType.DELETE:
      axios
        .delete(`/api/todo/${action.payload}`)
        .then((response) => console.log(response));
      return todos.filter((todo) => todo.id !== action.payload);
    case ActionsType.TOGGLE:
      return todos.map((todo) => {
        if (todo.id === action.payload) {
          let data = { ...todo, isDone: !todo.isDone };
          axios
            .put(`/api/todo/update/${todo.id}`, data)
            .then((response) => console.log(response));
          return data;
        } else return todo;
      });
    case ActionsType.UPDATE:
      return todos.map((todo) => {
        if (todo.id === action.payload.id) {
          let data = { ...todo, todo: action.payload.todo };
          axios
            .put(`/api/todo/update/${todo.id}`, data)
            .then((response) => console.log(response));
          return data;
        } else return todo;
      });
    default:
      return todos;
  }
};

const App: React.FC = () => {
  const [todos, todoDispatch] = useReducer(TodoReducer, []);
  const [loading, setLoading] = useState(true);

  const onDragEnd = (result: DropResult) => {
    const { draggableId, source, destination } = result;
    if (!destination) return;
    if (
      destination.droppableId === source.droppableId &&
      destination.index === source.index
    )
      return;

    if (destination.droppableId !== source.droppableId) {
      todoDispatch({ type: ActionsType.TOGGLE, payload: +draggableId });
    }
  };

  useEffect(() => {
    axios.get("/api/todo").then((response) => {
      todoDispatch({ type: ActionsType.GET, payload: response.data });
      setLoading(false);
    });
  }, []);

  return loading ? (
    <h1>Loading...</h1>
  ) : (
    <DragDropContext onDragEnd={onDragEnd}>
      <div className="App">
        <span className="heading">Taskify</span>
        <InputField dispatch={todoDispatch} />
        <TodoList todos={todos} dispatch={todoDispatch} />
      </div>
    </DragDropContext>
  );
};

export default App;
