import React from "react";
import { Link } from "react-router-dom";

function Login() {
  return (
    <div className="container">
      <div className="form-container">
        <h2>Login </h2>
        <form className="user-form">
          <input type="text" name="username" placeholder="Username" />
          <input type="password" name="password" placeholder="Password" />
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
