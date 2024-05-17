import { redirect } from "next/navigation";
import { checkCookie, deleteCookie } from "./components/api";
import NavBar from "./header/page";

export default async function Home() {
  const isLoggedIn = await checkCookie();
  if (!isLoggedIn) {
    redirect("/login");
  }

  return (
    <div>
      <NavBar />

      <div className="container-fluid">
        <h3>Hello</h3>
      </div>
    </div>
  );
}
