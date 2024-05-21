import React, { useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import { useNavigate } from "react-router-dom";

function Login() {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });
  const navigate = useNavigate();

  const handleFormChange = (e) => {
    setFormData((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  const handleLogin = async (e) => {
    // prevent default
    e.preventDefault();
    if (formData.username.length < 1 || formData.password < 1)
      alert("All fields are mandatory");

    let tosendData = {
      username: formData.username,
      password: formData.password,
    };

    console.log("FormData ", tosendData);

    let res = await axios.post("http://localhost:9090/users/login", tosendData);

    console.log("LOGIN RESPONSE ", res);

    localStorage.setItem("todos-username", tosendData.username); // will be removed

    // TODO: save the user details in cookie session

    navigate("/");
  };

  return (
    <div className="container">
      <div className="form-container">
        <h2>Login </h2>
        <form className="user-form" onSubmit={handleLogin}>
          <input
            type="text"
            name="username"
            placeholder="Username"
            onChange={handleFormChange}
          />
          <input
            type="password"
            name="password"
            placeholder="Password"
            onChange={handleFormChange}
          />
          <input type="submit" value="Login" />
        </form>
        <p>
          Don't have an account?{" "}
          <span>
            <Link to={`/signup`}>Sign Up</Link>
          </span>
        </p>
      </div>
    </div>
  );
}

export default Login;
