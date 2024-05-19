import React, { useEffect, useState } from "react";
import axios from "axios";
import TodoItem from "./components/TodoItem";

//styles
import "./index.css";
import "react-calendar/dist/Calendar.css";
import CreateTodo from "./components/CreateTodo";
import Login from "./Login";
import Signup from "./Signup";

const App = () => {
  const [todos, setTodos] = useState();
  const fetchTodos = async () => {
    let res = await axios.get("http://localhost:9090");
    setTodos(res.data);
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  const onDelete = async (id) => {
    window.location.reload(true);
    await axios.delete(`http://localhost:9090/${id}`);
  };

  const onUpdate = async (id, data) => {
    await axios.put(`http://localhost:9090/${id}`, data);
  };

  return (
    <div className="todos-container">
      <h2>Todo Application</h2>
      <div className="todos-list">
        {todos &&
          todos.map((t) => (
            <TodoItem
              todo={t}
              key={t.id}
              onDelete={onDelete}
              onUpdate={onUpdate}
            />
          ))}
        <CreateTodo />
      </div>
    </div>
  );
};

export default App;
