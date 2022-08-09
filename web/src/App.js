import { useRef, useState } from 'react';
import '@picocss/pico';
import './App.css';

const Signin = () => {
  const name = useRef(null);
  const username = useRef(null);
  const bio = useRef(null);
  const [isBusy, setBusy] = useState(false);
  const [error, setError] = useState("");
  const signin = async () => {
    if (isBusy) {
      return;
    } else if (name.current.value.trim() === "") {
      setError("Your name can't be empty!");
      return;
    } else {
      setBusy(true);
      setError("");
    }
    let data = { user: name.current.value.trim() };
    if (username.current.value.trim() !== "") {
      data.username = username.current.value.trim();
    }
    if (bio.current.value.trim() !== "") {
      data.bio = bio.current.value.trim();
    }
    const response = await fetch(
      "/api/v1/createUser",
      {
        method: 'POST',
        body: JSON.stringify(data),
        headers: { 'Content-type': 'application/json; charset=UTF-8' },
      },
    );
    setBusy(false); // Slow jobs already done, all jobs now run too fast
    if(!response.ok && response.status !== 400) {
      return setError("Failed to send request!");
    }
    const respJSON = await response.json();
    if(respJSON.error) {
      return setError(respJSON.error)
    }
    localStorage.setItem("token", respJSON.token);
  };
  return (
    <div className="container">
      <h1>Welcome!</h1>
      <div className="grid">
        <label htmlFor="name">
          Name:
          <input ref={name} name="name" placeholder="Your name..." />
        </label>
        <label htmlFor="username">
          Username:
          <input ref={username} name="username" placeholder="A new username..." />
        </label>
        <label htmlFor="bio">
          Bio:
          <textarea ref={bio} name="bio" placeholder="Some information about you..." />
        </label>
        <span className="error">{error}</span>
        <button onClick={signin} aria-busy={isBusy}>Sign-in</button>
      </div>
      <span>Already have an account? <a href="/login">Login</a></span>
    </div>
  );
}

const App = () => (
  <div className="container-fluid">
    <Signin />
  </div>
);

export default App;
