// import { useState, useEffect } from 'react';
// // for use in code
// const useRoles = () => {
//   const [userRoles, setUserRoles] = useState([]);

//   useEffect(() => {
//     const roles = JSON.parse(sessionStorage.getItem('userRoles') || '[]');
//     setUserRoles(roles);
//   }, []);

//   const hasRole = (role) => userRoles.includes(role);

//   return { userRoles, hasRole };
// };

// export default useRoles;

import { useState, useEffect } from 'react';
import { useAuth0 } from "@auth0/auth0-react";
// for use in code
const useRoles = () => {
  const [userRoles, setUserRoles] = useState([]);
  const { loginWithRedirect, logout, user, isLoading } = useAuth0();

  useEffect(() => {
    // const roles = JSON.parse(sessionStorage.getItem('userRoles') || '[]');
    const roles = user?.user_metadata?.roles || '[]'
    setUserRoles(roles);
  }, []);

  const hasRole = (role) => userRoles.includes(role);

  return { userRoles, hasRole };
};

export default useRoles;
