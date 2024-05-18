"use server";
import { cookies } from "next/headers";

export async function createCookie(data: string) {
  cookies().set("session", data, { sameSite: "none", secure: true });
}

export async function checkCookie() {
  const session = cookies().get("session")?.value;

  if (session == undefined) {
    return false;
  } else return true;
}

export async function deleteCookie() {
  cookies().delete("session");
}

// export default async function getUser() {
//   const response = await fetch("http://localhost8000/api/users/test");

//   if (!response.ok) {
//     throw new Error("Failed to fetch data");
//   }

//   return response.json();
// }
