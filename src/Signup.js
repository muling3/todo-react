import React from "react";
import { Link } from "react-router-dom";

function Signup() {
  return (
    <div className="container">
      <div className="form-container">
        <h2>Sign Up </h2>
        <form className="user-form">
          <input type="text" name="username" placeholder="Username" />
          <input type="password" name="password" placeholder="Password" />
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
