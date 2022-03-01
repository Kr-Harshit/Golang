import React from "react";
import "./TodoList.css";
import { Todo, Actions } from "../model";
import TodoCard from "./TodoCard";
import { Droppable } from "react-beautiful-dnd";

interface Props {
	todos: Todo[];
	dispatch: React.Dispatch<Actions>;
}

const TodoList: React.FC<Props> = ({ todos, dispatch }) => {
	return (
		<div className="todo__container">
			<Droppable droppableId="activeTodos">
				{(provided) => (
					<div
						className="todoList active"
						ref={provided.innerRef}
						{...provided.droppableProps}
					>
						<span className="todos__heading">Active Task</span>
						{todos.map((todo, idx) =>
							!todo.isDone ? (
								<TodoCard
									key={idx}
									index={idx}
									id={todo.id}
									todo={todo.todo}
									deadline={todo.deadline}
									isDone={todo.isDone}
									dispatch={dispatch}
								/>
							) : (
								<div key={idx}></div>
							)
						)}
						{provided.placeholder}
					</div>
				)}
			</Droppable>
			<Droppable droppableId="completedTodos">
				{(provided) => (
					<div
						className="todoList completed"
						ref={provided.innerRef}
						{...provided.droppableProps}
					>
						<span className="todos__heading">Completed Task</span>
						{todos.map((todo, idx) =>
							todo.isDone ? (
								<TodoCard
									key={idx}
									index={idx}
									id={todo.id}
									todo={todo.todo}
									deadline={todo.deadline}
									isDone={todo.isDone}
									dispatch={dispatch}
								/>
							) : (
								<div key={idx}></div>
							)
						)}
						{provided.placeholder}
					</div>
				)}
			</Droppable>
		</div>
	);
};

export default TodoList;
