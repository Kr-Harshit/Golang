import React, { useEffect, useState, useRef } from "react";
import { ActionsType, Actions } from "../model";
import {
	EditTwoTone,
	DeleteFilled,
	CheckCircleTwoTone,
	CloseCircleTwoTone,
} from "@ant-design/icons";
import "./TodoCard.css";
import { Draggable } from "react-beautiful-dnd";

interface Props {
	id: number;
	index: number;
	todo: string;
	deadline: Date;
	isDone: boolean;
	dispatch: React.Dispatch<Actions>;
}

function getDeadlineString(seconds: number): string {
	let d = Math.floor(seconds / (3600 * 24)),
		h = Math.floor((seconds % (3600 * 24)) / 3600),
		m = Math.floor((seconds % 3600) / 60),
		s = Math.floor(seconds % 60);

	let dDisplay = d > 0 ? d + (d === 1 ? "day, " : "days, ") : "",
		hDisplay = h > 0 ? h + "hr  " : "",
		mDisplay = m > 0 ? m + "min " : "",
		sDisplay = s > 0 ? s + "sec " : "";

	return dDisplay + hDisplay + mDisplay + sDisplay;
}

const TodoCard: React.FC<Props> = ({
	id,
	index,
	todo,
	deadline,
	isDone,
	dispatch,
}) => {
	const [timer, setTimer] = useState<number>(
		Math.abs(deadline.getTime() - new Date().getTime()) / 1000
	);
	const [task, setTask] = useState<string>(todo);
	const [edit, setEdit] = useState<boolean>(false);

	const isDoneHandler = (id: number) => {
		dispatch({ type: ActionsType.TOGGLE, payload: id });
	};

	const deleteTodoHandler = (id: number) => {
		dispatch({ type: ActionsType.DELETE, payload: id });
	};

	const updateTodoHandler = (
		e: React.FormEvent<HTMLFormElement>,
		id: number
	) => {
		e.preventDefault();
		dispatch({ type: ActionsType.UPDATE, payload: { id: id, todo: task } });
		setEdit(false);
	};

	const inputRef = useRef<HTMLInputElement>(null);

	useEffect(() => {
		if (timer > 0)
			setInterval(() => {
				setTimer((counter) => counter - 1);
			}, 1000);
		return () => {
			setTimer(0);
			setTask("");
			setEdit(false);
		};
	}, [isDone]);

	useEffect(() => {
		return () => {
			setTimer(0);
			setTask("");
			setEdit(false);
		};
	}, []);

	useEffect(() => {
		inputRef.current?.focus();
	}, [edit]);

	return (
		<Draggable draggableId={id.toString()} index={index}>
			{(provided) => (
				<form
					className="todo__card"
					onSubmit={(e) => updateTodoHandler(e, id)}
					{...provided.draggableProps}
					{...provided.dragHandleProps}
					ref={provided.innerRef}
				>
					<div className="todo__content">
						<input
							type="text"
							ref={inputRef}
							value={task}
							className="todo__text"
							disabled={!edit}
							onChange={(e) => setTask(e.target.value)}
						/>
						<span className="todo__deadline">{getDeadlineString(timer)}</span>
					</div>
					<div className="todo__control">
						<EditTwoTone
							className="todo__icon"
							onClick={() => {
								if (!isDone && !edit) {
									setEdit(!edit);
								}
								inputRef.current?.focus();
							}}
						/>
						<DeleteFilled
							className="todo__icon"
							style={{ color: "rgba(0,0,0,0.7)" }}
							onClick={() => deleteTodoHandler(id)}
						/>
						{!isDone ? (
							<CheckCircleTwoTone
								className="todo__icon"
								twoToneColor="#52c41a"
								onClick={() => isDoneHandler(id)}
							/>
						) : (
							<CloseCircleTwoTone
								className="todo__icon"
								twoToneColor="#eb2f96"
								onClick={() => isDoneHandler(id)}
							/>
						)}
					</div>
				</form>
			)}
		</Draggable>
	);
};

export default TodoCard;
