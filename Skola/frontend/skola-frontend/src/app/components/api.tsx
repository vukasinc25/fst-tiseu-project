"use server";
import { cookies } from "next/headers";

export async function getDiplomasByUser(userId: any) {
  const response = await fetch("/http://localhost8000/api/skola/diplomas", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(userId),
  });
}

export async function createCookie(data: string) {
  cookies().set("session", data, { sameSite: "none", secure: true });
}

export async function checkCookie() {
  const session = cookies().get("session")?.value;

  return session != undefined;
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
