import { useRef, useState } from "react";
import "./index.css";

const Login = () => {
  const token = useRef(null);
  const [isBusy, setBusy] = useState(false);
  const login = async () => {
    if(isBusy) {
      return;
    }
    setBusy(true);
  };
  return (
    <div className="container conteiner-size">
      <h1>Welcome back!</h1>
      <div className="container">
        <label htmlFor="token">
          Token:
          <input ref={token} placeholder="Your login token..." />
        </label>
        <button onClick={login} aria-busy={isBusy}>Login</button>
      </div>
      Do not have an account? <a href="/signin">Create one</a>
    </div>
  );
}

export default Login;