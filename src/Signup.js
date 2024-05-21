import React, { useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import { useNavigate } from "react-router-dom";

function Signup() {
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

  const handleSignup = async (e) => {
    // prevent default
    e.preventDefault();
    if (formData.username.length < 1 || formData.password < 1)
      alert("All fields are mandatory");

    let tosendData = {
      username: formData.username,
      password: formData.password,
    };

    console.log("Signup FormData ", tosendData);

    let res = await axios.post("http://localhost:9090/users/", tosendData);

    console.log("SIGNUP RESPONSE ", res);

    navigate("/login");
  };

  return (
    <div className="container">
      <div className="form-container">
        <h2>Sign Up </h2>
        <form className="user-form" onSubmit={handleSignup}>
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
          <input type="submit" value="Sign Up" />
        </form>
        <p>
          Already have an account?{" "}
          <span>
            <Link to={`/login`}>Login</Link>
          </span>
        </p>
      </div>
    </div>
  );
}

export default Signup;
