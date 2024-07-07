"use server";
import { error } from "console";
import { cookies } from "next/headers";

export async function getDiplomasByUser(userId: any) {
  try {
    const response = await fetch("http://localhost:8002/skola/diplomas", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ userId: userId }),
    });
  } catch (error) {
    console.log(error);
  }
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
