"use client";

import { useRouter } from "next/navigation";
import { deleteCookie } from "../components/api";

export default function NavBar() {
  const router = useRouter();

  function handleClick() {
    deleteCookie();
    router.push("/login");
  }

  return (
    <div className="container-fluid mb-3">
      <nav className="navbar navbar-expand-lg bg-primary" data-bs-theme="dark">
        <div className="container-fluid">
          <a className="navbar-brand">Skola</a>
          <button
            className="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarNavAltMarkup"
            aria-controls="navbarNavAltMarkup"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarNavAltMarkup">
            <div className="navbar-nav">
              <a className="nav-link">Home</a>
              <a className="nav-link">Profile</a>
              <button className="nav-link" onClick={handleClick}>
                Logout
              </button>
            </div>
          </div>
        </div>
      </nav>
    </div>
  );
}
