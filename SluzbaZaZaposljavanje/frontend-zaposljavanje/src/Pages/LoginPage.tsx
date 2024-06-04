import React, { useState } from 'react';
import "./LoginPage.css"
import { useNavigate } from 'react-router';

function LoginPage() {
  const [username, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [jwtToken, setToken] = useState("");
  let navigate = useNavigate()

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      const response = await fetch("http://localhost:8000/users/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
        
      });
  
      if (response.ok) {
        const data = await response.json();
        setToken(data)
        localStorage.setItem("jwtToken", data.access_token)
        console.log("Login successful");
        return navigate("/")
      } else {
        console.error("Login failed");
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <div>
      <section className="vh-100 gradient-custom">
        <div className="container py-5 h-100">
          <div className="row d-flex justify-content-center align-items-center h-100">
            <div className="col-12 col-md-8 col-lg-6 col-xl-8">
              <div className="card bg-dark text-white" style={{ borderRadius: "1rem" }}>
                <div className="card-body p-5 text-center">

                  <div className="mb-md-5 mt-md-4 pb-5">
                    <h2 className="fw-bold mb-2 text-uppercase">Login</h2>
                    <p className="text-white-50 mb-5">Please enter your username and password!</p>

                    <form onSubmit={handleSubmit}>
                      <div className="form-outline form-white mb-4">
                        <input type="text" id="typeEmailX-2" className="form-control form-control-lg" value={username} onChange={(e) => setEmail(e.target.value)} />
                        <label className="form-label" htmlFor="typeEmailX-2">Username</label>
                      </div>

                      <div className="form-outline form-white mb-4">
                        <input type="password" id="typePasswordX-2" className="form-control form-control-lg" value={password} onChange={(e) => setPassword(e.target.value)} />
                        <label className="form-label" htmlFor="typePasswordX-2">Password</label>
                      </div>

                      <p className="small mb-5 pb-lg-2"><a className="text-white-50" href="#!">Forgot password?</a></p>

                      <button className="btn btn-outline-light btn-lg px-5" type="submit">Login</button>
                    </form>

                  </div>

                  <div>
                    <p className="mb-0">Don't have an account? <a href="#!" className="text-white-50 fw-bold">Sign Up</a></p>
                  </div>

                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}

export default LoginPage;
