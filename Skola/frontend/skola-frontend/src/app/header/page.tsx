"use client";

import { useRouter } from "next/navigation";
import { deleteCookie } from "../components/api";
import Link from "next/link";
import { useUser } from "@auth0/nextjs-auth0/client";

export default function NavBar() {
  const router = useRouter();
  const { user, error, isLoading } = useUser();

  // function handleClick() {
  //   deleteCookie();
  //   router.push("/login");
  // }

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
              {/* <button className="nav-link" onClick={handleClick}>
                Logout
              </button> */}
              {user ? (
                <a className="nav-link" href="/api/auth/logout">
                  Logout
                </a>
              ) : (
                <a className="nav-link" href="/api/auth/login">
                  Login
                </a>
              )}
              <Link href="/api/auth/logout"></Link>
            </div>
          </div>
        </div>
      </nav>
    </div>
  );
}
