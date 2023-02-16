import React, { useEffect, useState } from "react";
import { format,add } from "date-fns";

//icons
import { CaretDownOutlined, CaretUpOutlined, RightCircleOutlined, RotateLeftOutlined } from '@ant-design/icons'

const TodoItem = ({ todo, onDelete, onUpdate }) => {
  const [open, setOpen] = useState(false);
  const [formData, setFormData] = useState({
    body: todo.body,
    priority: todo.priority,
  });

  const handleExpandTodo = (e) => {
    setOpen((prev) => !prev);
  };

  const handleFormChange = (e) => {
    setFormData(prev => ({...prev, [e.target.name]:e.target.value}))
  };

  const handleUpdateTodo = () =>{
    let tosendData = {
      body: formData.body,
      priority: formData.priority,
    };

    onUpdate(todo.id, tosendData)
  }
  return (
    <div className="item" style={{ height: open ? "300px" : "50px" }}>
      <div className="top" onClick={handleExpandTodo}>
        <div className="title">
          <div className="icon">
            <RightCircleOutlined />
          </div>
          <div className="icon"></div>
          <p>{todo.title}</p>
        </div>
        <div className="due-date">
          <p>{format(new Date(todo.due_date.Time), "dd/MM/yyyy")}</p>
          <div className="icon" onClick={handleExpandTodo} id="caret" style={{ transform: open ? "rotate(-180deg)" : "rotate(0deg)"}}>
             <CaretDownOutlined />
          </div>
        </div>
      </div>
      <hr style={{ opacity: open ? 1 : 0 }} />
      <div className="body">
        <div className="notes-container">
          <p>Notes</p>
          <div className="div">
            <textarea
              name="body"
              id=""
              cols="30"
              rows="10"
              defaultValue={todo.body}
              onChange={handleFormChange}
            ></textarea>
          </div>
        </div>
        <div className="date-container">
          <div className="dates">
            <div className="date-item">
              <p>Due date</p>
              <div className="div">
                <input
                  name="due_date"
                  id="date"
                  type="date"
                  min={format(new Date(), "yyyy-MM-dd")}
                  max={format(add(new Date(), { days: 30 }), "yyyy-MM-dd")}
                  defaultValue={format(
                    new Date(todo.due_date.Time),
                    "yyyy-MM-dd"
                  )}
                  readOnly={true}
                ></input>
              </div>
            </div>
            <div className="date-item">
              <p>Priority</p>
              <div className="div">
                <select
                  name="priority"
                  id="priority"
                  defaultValue={todo.priority}
                  onChange={handleFormChange}
                >
                  <option value="LOW">Low</option>
                  <option value="MEDIUM">Medium</option>
                  <option value="HIGH">High</option>
                </select>
              </div>
            </div>
          </div>
          <div className="footer">
            <button onClick={handleUpdateTodo}>Update</button>
            <button onClick={() => onDelete(todo.id)}>Delete</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default TodoItem;
