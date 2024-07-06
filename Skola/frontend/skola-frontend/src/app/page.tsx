"use client";
import { redirect } from "next/navigation";
import { checkCookie, deleteCookie, getDiplomasByUser } from "./components/api";
import NavBar from "./header/page";
import { useUser } from "@auth0/nextjs-auth0/client";
import { useRouter } from "next/navigation";

export default function Home() {
  const { user, error, isLoading } = useUser();
  const router = useRouter();
  if (isLoading)
    return <div className="mx-auto text-center fs-3 loading ">Loading...</div>;
  // if (user == undefined) {
  //   router.push("/api/auth/login");
  // }

  // console.log(user?.user_metadata);
  console.log(user);
  const userId = user?.user_metadata;
  // getDiplomasByUser(userId);
  // const isLoggedIn = await checkCookie();
  // if (!isLoggedIn) {
  //   redirect("/login");
  // }

  return (
    <div>
      <NavBar />
      <div className="container-fluid">
        <h2>year 3</h2>
        <div className="row text-center">
          <div className="col">
            <ul className="list-group list-group-flush">
              <li className="list-group-item">Srpski</li>
              <li className="list-group-item">Matematika</li>
              <li className="list-group-item">Fizicko</li>
              <li className="list-group-item">Fizika</li>
              <li className="list-group-item">Hemija</li>
            </ul>
          </div>
          <div className="col">
            <ul className="list-group list-group-flush">
              <li className="list-group-item">5, 5, 5, 5, 5, 5, 5, 5 | 5</li>
              <li className="list-group-item">4, 5, 2, 3, 4, 3, 4, 5 | 4,5</li>
              <li className="list-group-item">ocene 3</li>
              <li className="list-group-item">ocene 4</li>
              <li className="list-group-item">ocene 5</li>
            </ul>
          </div>
        </div>

        <h2>year 2</h2>
        <div className="row text-center">
          <div className="col">
            <ul className="list-group list-group-flush">
              <li className="list-group-item">Srpski</li>
              <li className="list-group-item">Matematika</li>
              <li className="list-group-item">Fizicko</li>
              <li className="list-group-item">Fizika</li>
              <li className="list-group-item">Hemija</li>
            </ul>
          </div>
          <div className="col">
            <ul className="list-group list-group-flush">
              <li className="list-group-item">5, 5, 5, 5, 5, 5, 5, 5 | 5</li>
              <li className="list-group-item">4, 5, 2, 3, 4, 3, 4, 5 | 4,5</li>
              <li className="list-group-item">ocene 3</li>
              <li className="list-group-item">ocene 4</li>
              <li className="list-group-item">ocene 5</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}
