// src/utils/jwt.ts
import { jwtDecode } from 'jwt-decode';

interface UserInfo {
  username: string;
  issued_at: string;
  role: string;
  expired_at: string;
}

export const decodeToken = (token: string): UserInfo | null => {
  try {
    const decoded: UserInfo = jwtDecode(token);
    return decoded;
  } catch (error) {
    console.error('Invalid token:', error);
    return null;
  }
};
