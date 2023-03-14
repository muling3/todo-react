
import axios from "axios";
import { differenceInCalendarDays, sub } from "date-fns";
import { format, add } from "date-fns/esm";
import React, { useRef, useState } from "react";  

//icons
import { CaretDownOutlined, PlusOutlined } from '@ant-design/icons'

const CreateTodo = () => {
  const inputRef = useRef();
  const [open, setOpen] = useState(false);
  const [formData, setFormData] = useState({
    title: '',
    body: '',
    due_date: '',
    priority: ''
  });

  const handleExpandTodo = (e) => {
    setOpen((prev) => !prev);
  };

  const handleEndEditing = (e) => {
    if (e.key === "Enter") {
      setOpen(true)
    }
  };

  const handleFormChange = (e) => {
    setFormData((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  const handleSaveTodo = async() => {
    let dateDiff = differenceInCalendarDays(
      new Date(formData.due_date),
      new Date()
    );
    let tosendData = {
      title: formData.title,
      body: formData.body,
      due: dateDiff ? dateDiff : 0,
      priority: formData.priority ? formData.priority : "LOW",
    };

    let res = await axios.post("http://localhost:9090", tosendData);
    setOpen(prev => !prev)
    window.location.reload(true)
  };

  return (
    <div className="item" style={{ height: open ? "300px" : "50px" }}>
      <div className="top" onClick={() => inputRef.current.focus()}>
        <div className="title">
          <div className="icon">
            <PlusOutlined />
          </div>
          <input
            ref={inputRef}
            type="text"
            name="title"
            id="new"
            placeholder="New Task"
            autoComplete="off"
            onKeyUp={handleEndEditing}
            onChange={handleFormChange}
          />
        </div>
        <div className="due-date">
          <div className="icon" onClick={handleExpandTodo} id="caret" style={{ transform: open ? "rotate(-180deg)" : "rotate(0deg)"}}>
            <CaretDownOutlined />
          </div>
        </div>
      </div>
      <hr style={{ opacity: open ? 1 : 0 }} />
      <div
        className="body"
        style={{ opacity: open ? 1 : 0, display: open ? "grid" : "none" }}
      >
        <div className="notes-container">
          <p>Notes</p>
          <div className="div">
            <textarea
              name="body"
              id=""
              cols="30"
              rows="10"
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
                  defaultValue={format(new Date(), "yyyy-MM-dd")}
                  onChange={handleFormChange}
                ></input>
              </div>
            </div>
            <div className="date-item">
              <p>Priority</p>
              <div className="div">
                <select
                  name="priority"
                  id="priority"
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
            <button onClick={handleSaveTodo}>Save</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CreateTodo;
