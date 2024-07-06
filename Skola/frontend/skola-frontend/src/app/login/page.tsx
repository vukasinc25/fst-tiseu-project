"use client";

import { ChangeEvent, FormEvent, useState } from "react";
import { createCookie } from "../components/api";
import { useRouter } from "next/navigation";

export default function Login() {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });

  const router = useRouter();

  function handleChange(event: ChangeEvent<HTMLInputElement>) {
    // console.log(event.target.name + " " + event.target.value);
    // console.log(formData);
    setFormData((prev) => {
      return { ...prev, [event.target.name]: event.target.value };
    });
  }

  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const response = await fetch("http://localhost:8000/users/login", {
      method: "POST",
      body: JSON.stringify(formData),
    });

    const data = await response.json();
    console.log(data);
    await createCookie(data.access_token);

    localStorage.setItem("jwt", data);

    router.push("/");
  }

  return (
    <div className="d-flex justify-content-center" style={{ marginTop: 200 }}>
      <form onSubmit={onSubmit}>
        <legend>Login</legend>
        <div className="form-group">
          <input
            type="text"
            className="form-control mb-2"
            id="username"
            name="username"
            placeholder="username"
            onChange={handleChange}
            value={formData.username}
          />
        </div>

        <div className="form-group">
          <input
            type="password"
            className="form-control"
            id="password"
            name="password"
            placeholder="Password"
            onChange={handleChange}
            value={formData.password}
          />
        </div>

        <div className="d-flex justify-content-end">
          <button
            type="submit"
            className="btn btn-primary mt-3"
            disabled={formData.username == "" || formData.password == ""}
          >
            Login
          </button>
        </div>
      </form>
    </div>
  );
}
