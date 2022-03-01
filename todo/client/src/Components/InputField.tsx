import React, { useState } from "react";
import "./InputField.css";
import { ActionsType, Actions } from "../model";

interface Props {
	dispatch: React.Dispatch<Actions>;
}

function getTimeString(datetime: Date): string {
	let month = `0${datetime.getMonth() + 1}`,
		year = String(datetime.getFullYear()),
		date =
			datetime.getDate() > 9
				? String(datetime.getDate())
				: `0${datetime.getDate()}`,
		hour =
			datetime.getHours() > 9
				? String(datetime.getHours())
				: `0${datetime.getHours()}`,
		minute =
			datetime.getMinutes() > 9
				? String(datetime.getMinutes())
				: `0${datetime.getMinutes()}`;
	let timeString = `${year}-${month}-${date}T${hour}:${minute}`;
	return timeString;
}

const InputField: React.FC<Props> = ({ dispatch }) => {
	const [todo, setTodo] = useState<string>("");
	const [deadline, setDeadline] = useState<Date>(new Date());

	const addTodoHandler = (e: React.FormEvent) => {
		e.preventDefault();
		dispatch({
			type: ActionsType.ADD,
			payload: { todo: todo, deadline: deadline },
		});
		setTodo("");
		setDeadline(new Date());
	};

	
	return (
		<form className="input" onSubmit={(e) => addTodoHandler(e)}>
			<div className="flex__wrapper">
				<input
					type="text"
					placeholder="Enter a Todo"
					id="input__text"
					className="input__box"
					onChange={(e) => setTodo(e.target.value)}
					value={todo}
					required
				/>
			</div>
			<div className="flex__wrapper">
				<div className="input__date__wrapper">
					<label htmlFor="input__date" className="input__label">
						<span>Task Deadline</span>
					</label>
					<div className="deadline__inner__wrapper">
						<input
							type="datetime-local"
							className="input__box"
							id="input__date"
							min={getTimeString(new Date())}
							value={getTimeString(deadline)}
							onChange={(e) => setDeadline(new Date(e.target.value))}
						/>
						<span
							className="deadline__rest"
							onClick={() => setDeadline(new Date())}
						>
							reset
						</span>
					</div>
				</div>
				<button
					type="submit"
					className="input__submit"
					disabled={!Boolean(todo)}
				>
					Add Task
				</button>
			</div>
		</form>
	);
};

export default InputField;
